#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;
use POSIX qw/ceil/;
use POSIX qw/floor/;

my $data_type = "";
if (1 <= @ARGV){
  $data_type = lc($ARGV[0]);
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

my $tbl_name_prefix = "test";
my @data_values = (-1.2, 2.4, 3.6, -4.2, 5.8, -6.2, 7.0, 8.1, -9.8, 10.1);

my $cnt = @data_values;
my $min = $data_values[0];
my $max = $data_values[$cnt - 1];
my $sum = 0;
my $avg = 0;
for my $value (@data_values) {
    $sum += $value;
}
$avg = $sum / @data_values;

my $tbl_name = $tbl_name_prefix . "_" . lc($data_type);

my $column_type = uc($data_type);
my $column_name = "c" . lc($data_type);
print "CREATE TABLE ${tbl_name} (\n";  
print "\t$column_name $column_type PRIMARY KEY";  
print "\n);\n";
print "{\n";
print "}\n";

for (my $n = 0; $n < @data_values; $n++){
  my $data_value = $data_values[$n];
  if ($data_type eq "int") {
    $data_value = int($data_value);
  }
  print "INSERT INTO ${tbl_name} ($column_name) VALUES ($data_value);\n";  
  print "{\n";  
  print "}\n";
  my @sels = ("ABS($column_name)", "FLOOR($column_name)", "CEIL($column_name)");
  my @expecteds = (abs($data_value), floor($data_value), ceil($data_value));
  for (my $n = 0; $n < @sels; $n++){
    my $sel = $sels[$n];
    my $exp = $expecteds[$n];
    print "SELECT $sel FROM ${tbl_name} WHERE $column_name = $data_value;\n";  
    print "{\n"; 
    print "\t\"rows\" :\n";  
    print "\t[\n";
    print "\t\t{\n";
    print "\t\t\t\"$sel\" : $exp\n";
    print "\t\t}\n";
    print "\t]\n";
    print "}\n";
  }
  print "DELETE FROM ${tbl_name} WHERE $column_name = $data_value;\n";
  print "{\n";  
  print "}\n";
}
