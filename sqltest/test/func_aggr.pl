#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

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

my @data_types = ("INT", "FLOAT", "DOUBLE");
my @data_values = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

my $tbl_name = "test";

for (my $t = 0; $t < @data_types; $t++){
  my $data_type = $data_types[$t];
  my $column_type = uc($data_type);
  my $column_name = "c" . lc($data_type);
  print "CREATE TABLE ${tbl_name} (\n";  
  print "\t$column_name $column_type PRIMARY KEY";  
  print "\n);\n";  
  print "{\n";  
  print "}\n";

  for (my $n = 0; $n < @data_values; $n++){
    my $data_value = $data_values[$n];
    print "INSERT INTO ${tbl_name} ($column_name) VALUES ($data_value);\n";  
    print "{\n";  
    print "}\n";
  }
}
