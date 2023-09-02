#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

my $pr_key_type = "";
if (1 <= @ARGV){
  $pr_key_type = $ARGV[0];
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

my $pg_type_file = "data/data_type.pict";
open(IN, $pg_type_file) or die "Failed to open $pg_type_file: $!";
my $line_no = 0;
my $pr_key_idx = -1;
while(<IN>){
  $line_no++;
  chomp($_);
  my @row = split(/\t/, $_, -1);
  if ($line_no <= 1) {
    print "CREATE TABLE test (\n";  
    for (my $i = 0; $i < scalar(@row); $i++) {
      my $type_name = $row[$i];
      my $column_name = "col" . $type_name;
      if ($type_name eq $pr_key_type) {
        print "\t" . $column_name . " " . $type_name . " PRIMARY KEY\n";  
      } else {
        print "\t" . $column_name . " " . $type_name . "\n";  
      }
    }
    print ");\n";  
    print "{\n";  
    print "}\n";  
  }
}
close(IN);

print<<FOTTER;
}
FOTTER