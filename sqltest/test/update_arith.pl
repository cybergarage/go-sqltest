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
my @seed_values = (100, 200, 300, 400, 500, 600, 700, 800, 900, 1000);
my @data_values;
for (my $n = 0; $n < @seed_values; $n++){
  my $seed_value = $seed_values[$n];
  if ($data_type eq "int") {
    $seed_value = int($seed_value);
  }
  push(@data_values, $seed_value);
}

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
my $pk_column_name = "p" . lc($data_type);
my $v_column_name = "v" . lc($data_type);
print "CREATE TABLE ${tbl_name} (\n";  
print "\t$pk_column_name $column_type PRIMARY KEY, \n";  
print "\t$v_column_name $column_type";  
print "\n);\n";
print "{\n";
print "}\n";

for (my $n = 0; $n < @data_values; $n++){
  my $data_value = $data_values[$n];
  print "INSERT INTO ${tbl_name} ($pk_column_name, $v_column_name) VALUES ($data_value, $data_value);\n";  
  print "{\n";  
  print "}\n";
}

my @operators = ("+", "-", "*", "/", "%");
my @param_values = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
for (my $n = 0; $n < @operators; $n++){
  my $operator = $operators[$n];
  my $init_value = $seed_values[0];
  for (my $p = 0; $p < @param_values; $p++){
    my $param_value = $param_values[$p];
    my $where = "$pk_column_name = $init_value";
    my $update_value = $init_value;
    if ($operator eq "+") {
      $update_value = $init_value + $param_value;
    }
    if ($operator eq "-") {
      $update_value = $init_value - $param_value;
    }
    if ($operator eq "*") {
      $update_value = $init_value * $param_value;
    }
    if ($operator eq "/") {
      $update_value = $init_value / $param_value;
    }
    if ($operator eq "%") {
      $update_value = $init_value % $param_value;
    }
    my @set_values = ($init_value, "$v_column_name $operator $param_value");
    my @expected_values = ($init_value, $update_value);
    for (my $v = 0; $v < @set_values; $v++){
      my $set_value = $set_values[$v];
      my $expected_value = $expected_values[$v];
      print "UPDATE ${tbl_name} SET $v_column_name = $set_value WHERE $where;\n";  
      print "{\n";  
      print "}\n";
      print "SELECT $v_column_name FROM ${tbl_name} WHERE $where;\n";  
      print "{\n"; 
      print "\t\"rows\" :\n";  
      print "\t[\n";
      print "\t\t{\n";
      print "\t\t\t\"$v_column_name\" : $expected_value";
      print "\n";
      print "\t\t}\n";
      print "\t]\n";
      print "}\n";
    }
  }
}

for (my $n = 0; $n < @data_values; $n++){
  my $data_value = $data_values[$n];
  print "DELETE FROM ${tbl_name} WHERE $pk_column_name = $data_value;\n";
  print "{\n";  
  print "}\n";
}
