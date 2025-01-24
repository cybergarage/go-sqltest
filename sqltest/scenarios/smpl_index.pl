#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $idx_key_type = "";
if (1 <= @ARGV){
  $idx_key_type = lc($ARGV[0]);
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

my $idx_key_idx = -1;
my $pr_key_type = "";
my $pr_key_idx = -1;
for (my $n = 0; $n < scalar(@data_type_row); $n++) {
  my $type_name = lc($data_type_row[$n]);
  if ($type_name eq $idx_key_type) {
    $idx_key_idx = $n;
    $pr_key_idx = ($n + 1) % scalar(@data_type_row);
    $pr_key_type = lc($data_type_row[$pr_key_idx]);
    last;
  } 
}

if ($pr_key_idx < 0) {
  die "The primary key type ($pr_key_type) is not found in $data_type_file";
}

print "CREATE TABLE ${tbl_name} (\n";  
for (my $n = 0; $n < scalar(@data_type_row); $n++) {
  my $type_name = lc($data_type_row[$n]);
  my $column_type = uc($type_name);
  my $column_name = "c" . $type_name;
  if ($type_name eq $pr_key_type) {
    print "\t$column_name $column_type PRIMARY KEY";  
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

my $idx_column_name = "c" . $idx_key_type;
my $idx_name = $idx_key_type . "idx";
print "CREATE INDEX ${idx_name} ON ${tbl_name}(${idx_column_name});\n";  
print "{\n";  
print "}\n";

my @pr_data_array = ();
for my $data_row_no (0 .. $#data_rows) {
  my @data_row = @{$data_rows[$data_row_no]};
  my $pr_data = $data_row[$pr_key_idx];
  if (!grep { $_ eq $pr_data } @pr_data_array) {
    push(@pr_data_array, $pr_data);
  }
}

my @idx_data_array = ();
for my $data_row_no (0 .. $#data_rows) {
  my @data_row = @{$data_rows[$data_row_no]};
  my $idx_data = $data_row[$idx_key_idx];
  if (!grep { $_ eq $idx_data } @idx_data_array) {
    push(@idx_data_array, $idx_data);
  }
}

my @target_data_rows = (); 
for my $idx_data_no (0 .. $#idx_data_array) {
  my $idx_data = $idx_data_array[$idx_data_no];
  for my $data_row_no (0 .. $#data_rows) {
    my @data_row = @{$data_rows[$data_row_no]};
    if ($idx_data eq $data_row[$idx_key_idx]) {
      my $pr_data = $data_row[$pr_key_idx];
      if (grep { $_ eq $pr_data } @pr_data_array) {
        @pr_data_array = grep { $_ ne $pr_data } @pr_data_array;
        push(@target_data_rows, \@data_row);
      }
    }
  }
}

for my $idx_data_no (0 .. $#idx_data_array) {
  my $idx_data = $idx_data_array[$idx_data_no];
  my @inserted_data_rows = (); 

  my $where_column_name = "c" . lc($data_type_row[$idx_key_idx]);

  print "SELECT * FROM ${tbl_name} WHERE $where_column_name = ${idx_data};\n";  
  print "{\n";  
  print "}\n";  

  for my $data_row_no (0 .. $#target_data_rows) {
    my @data_row = @{$target_data_rows[$data_row_no]};
    if ($idx_data_array[$idx_data_no] ne $data_row[$idx_key_idx]) {
      next;
    }

    print "INSERT INTO ${tbl_name} (";  
    for (my $n = 0; $n < scalar(@data_row); $n++) {
      my $column_name = "c" . lc($data_type_row[$n]);
      if (0 < $n) {
        print ", ";
      }
      print $column_name;
    }
    print ") VALUES (";  
    for (my $n = 0; $n < scalar(@data_row); $n++) {
      if (0 < $n) {
        print ", ";  
      }
      print $data_row[$n];
    }
    print ");\n";  
    print "{\n";
    print "}\n";

    push(@inserted_data_rows, \@data_row);


    print "SELECT * FROM ${tbl_name} WHERE $where_column_name = ${idx_data};\n";  
    print "{\n";  
    print "\t\"rows\" :\n";  
    print "\t[\n";  
    for my $inserted_data_no (0 .. $#inserted_data_rows) {
      my @row = @{$inserted_data_rows[$inserted_data_no]};
      print "\t\t{\n";  
      for (my $n = 0; $n < scalar(@row); $n++) {
        my $column_name = "c" . lc($data_type_row[$n]);
        my $column_val = $row[$n];
        $column_val =~ s/'/"/g;
        print "\t\t\t\"$column_name\" : $column_val";
        if ($n < ((@row) - 1)) {
          print ",";  
        }
        print "\n";
      }
      print "\t\t}";  
      if ($inserted_data_no < ((@inserted_data_rows) - 1)) {
        print ",";  
      }
      print "\n";  
    }
    print "\t]\n";  
    print "}\n";  

  }

  print "DELETE FROM ${tbl_name} WHERE ${idx_column_name} = $idx_data_array[$idx_data_no];\n";  
}