#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $pr_key_type = "";
if (1 <= @ARGV){
  $pr_key_type = lc($ARGV[0]);
}

print<<HEADER;
- Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
-
- Licensed under the Apache License, Version 2.0 (the "License");
- you may not use this file except in compliance with the License.
- You may obtain a copy of the License at
-
-  http:-www.apache.org/licenses/LICENSE-2.0
-
- Unless required by applicable law or agreed to in writing, software
- distributed under the License is distributed on an "AS IS" BASIS,
- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
- See the License for the specific language governing permissions and
- limitations under the License.

HEADER

my $data_type_file = "data/data_type.pict";
my @data_type_row;
my @data_rows = ();
my $pr_key_idx = -1;
my $tbl_name = "test";
my $line_no = 0;

$line_no = 0;
open(IN, $data_type_file) or die "Failed to open $data_type_file: $!";
while(<IN>){
  $line_no++;
  chomp($_);
  my @row = split(/\t/, $_, -1);
  if ($line_no <= 1) {
    @data_type_row = @row;
  } else {
    $data_rows[$line_no - 2] = \@row;
  }
}
close(IN);

print "CREATE TABLE ${tbl_name} (\n";  
for (my $n = 0; $n < scalar(@data_type_row); $n++) {
  my $type_name = lc($data_type_row[$n]);
  my $column_type = uc($type_name);
  my $column_name = "c" . $type_name;
  if ($type_name eq $pr_key_type) {
    print "\t$column_name $column_type PRIMARY KEY";  
    $pr_key_idx = $n;
  } else {
    print "\t$column_name $column_type";  
  }
  if ($n < ((@data_type_row) - 1)) {
    print ",";  
  }
  print "\n";  
}
print ");\n";  
print "{\n";  
print "}\n";

if ($pr_key_idx < 0) {
  die "The primary key type ($pr_key_type) is not found in $data_type_file";
}

for my $row_no (0 .. $#data_rows) {
  my @row = @{$data_rows[$row_no]};

  print "INSERT INTO ${tbl_name} (";  
  for (my $n = 0; $n < scalar(@row); $n++) {
    my $type_name = lc($data_type_row[$n]);
    my $column_name = "c" . $type_name;
    if (0 < $n) {
      print ", ";
    }
    print $column_name;
  }
  print ") VALUES (";  
  for (my $n = 0; $n < scalar(@row); $n++) {
    if (0 < $n) {
      print ", ";  
    }
    print $row[$n];
  }
  print ");\n";  
  print "{\n";  
  print "}\n";

  my $type_name = lc($data_type_row[$pr_key_idx]);
  my $column_name = "c" . $type_name;

  print "SELECT * FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
  print "{\n";  
  print "\t\"rows\" :\n";  
  print "\t[\n";  
  print "\t\t{\n";  
  for (my $n = 0; $n < scalar(@row); $n++) {
    my $type_name = lc($data_type_row[$n]);
    my $column_name = "c" . $type_name;
    my $column_val = $row[$n];
    $column_val =~ s/'/"/g;
    print "\t\t\t\"$column_name\" : $column_val";
    if ($n < ((@row) - 1)) {
      print ",";  
    }
    print "\n";
  }
  print "\t\t}\n";  
  print "\t]\n";  
  print "}\n";  

  for my $update_row_no (0 .. $#data_rows) {
    if ($update_row_no == $row_no) {
      next;
    }
    my @update_row = @{$data_rows[$update_row_no]};

    print "UPDATE ${tbl_name} SET ";  
    my $n_colums = 0;
    for (my $n = 0; $n < scalar(@update_row); $n++) {
      my $type_name = lc($data_type_row[$n]);
      my $column_name = "c" . $type_name;
      if ($n == $pr_key_idx) {
        next;
      }
      if (0 < $n_colums) {
        print ", ";  
      }
      print "$column_name = $update_row[$n]";
      $n_colums++;
    }
    my $type_name = lc($data_type_row[$pr_key_idx]);
    my $column_name = "c" . $type_name;
    print " WHERE $column_name = $row[$pr_key_idx];\n";    
    print "{\n";  
    print "}\n";  

    my $type_name = lc($data_type_row[$pr_key_idx]);
    my $column_name = "c" . $type_name;
    print "SELECT * FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
    print "{\n";  
    print "\t\"rows\" :\n";  
    print "\t[\n";  
    print "\t\t{\n";  
    for (my $n = 0; $n < scalar(@row); $n++) {
      my $type_name = lc($data_type_row[$n]);
      my $column_name = "c" . $type_name;
      my $column_val;
      if ($n == $pr_key_idx) {
        $column_val = $row[$n];
      } else {
        $column_val = $update_row[$n];
      }
      $column_val =~ s/'/"/g;
      print "\t\t\t\"$column_name\" : $column_val";
      if ($n < ((@row) - 1)) {
        print ",";  
      }
      print "\n";
    }
    print "\t\t}\n";  
    print "\t]\n";  
    print "}\n";  
  }

  print "DELETE FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
  print "{\n";  
  print "}\n";  

  print "SELECT * FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
  print "{\n";  
  print "}\n";  
}
close(IN);

# $line_no = 0;
# open(IN, $data_type_file) or die "Failed to open $data_type_file: $!";
# while(<IN>){
#   $line_no++;
#   chomp($_);
#   my @row = split(/\t/, $_, -1);
#   if ($line_no <= 1) {
#     next;
#   }

#   print "UPDATE ${tbl_name} SET ";  
#   my $n_colums = 0;
#   for (my $n = 0; $n < scalar(@row); $n++) {
#     my $type_name = lc($data_type_row[$n]);
#     my $column_name = "c" . $type_name;
#     if ($n == $pr_key_idx) {
#       next;
#     }
#     if (0 < $n_colums) {
#       print ", ";  
#     }
#     print "$column_name = $row[$n]";
#     $n_colums++;
#   }
#   my $type_name = lc($data_type_row[$pr_key_idx]);
#   my $column_name = "c" . $type_name;
#   print " WHERE $column_name = $row[$pr_key_idx];\n";    
#   print "{\n";  
#   print "}\n";  

#   my $type_name = lc($data_type_row[$pr_key_idx]);
#   my $column_name = "c" . $type_name;
#   print "SELECT * FROM ${tbl_name} WHERE $column_name = $row[$pr_key_idx];\n";  
#   print "{\n";  
#   print "\t\"rows\" :\n";  
#   print "\t[\n";  
#   print "\t\t{\n";  
#   for (my $n = 0; $n < scalar(@row); $n++) {
#     my $type_name = lc($data_type_row[$n]);
#     my $column_name = "c" . $type_name;
#     my $column_val = $row[$n];
#     $column_val =~ s/'/"/g;
#     print "\t\t\t\"$column_name\" : $column_val";
#     if ($n < ((@row) - 1)) {
#       print ",";  
#     }
#     print "\n";
#   }
#   print "\t\t}\n";  
#   print "\t]\n";  
#   print "}\n";  
# }
# close(IN);
