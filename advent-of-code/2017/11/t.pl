#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;


my $dcoord = {
    n  => [-1, +0],
    ne => [-0.5, +1],
    se => [+0.5, +1],
    s  => [+1, +0],
    sw => [+0.5, -1],
    nw => [-0.5, -1],
};

sub solve_part_1($input_str) {
    my $dirs = parse_input($input_str);
    my ($y, $x) = walk($dirs);
    distance($y, $x);
}

sub solve_part_2($input_str) {
    my $dirs = parse_input($input_str);
    my $max_distance = 0;
    my ($y, $x) = (0, 0);
    for my $dir (@$dirs) {
        my $d = $dcoord->{$dir} // die "unknown dir $dir";
        $y += $d->[0];
        $x += $d->[1];
        my $distance = distance($y, $x);
        $max_distance = $distance if $distance > $max_distance;
    }
    $max_distance;
}

sub walk($dirs) {
    my ($y, $x) = (0, 0);
    for my $dir (@$dirs) {
        my $d = $dcoord->{$dir} // die "unknown dir $dir";
        $y += $d->[0];
        $x += $d->[1];
    }
    ($y, $x);
}

sub distance($y, $x) {
    abs($x) + abs($y) - abs($x)/2;
}

sub get_input($file) {
    open my $fh, '<', $file or die "$file: $!";
    my $line = <$fh>;
    chomp $line;
    close $fh;
    $line;
}

sub parse_input($input_str) {
    [ split /,/, $input_str ];
}


my $input_str = get_input('input.txt');

say "part 1: ", solve_part_1($input_str);
say "part 2: ", solve_part_2($input_str);
