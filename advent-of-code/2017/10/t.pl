#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

my $input = get_input();

say "part 1: ", solve_part_1($input);
say "part 2: ", solve_part_2($input);


sub solve_part_1($input) {
    my $lens = [ split /,/, $input ];
    my $list = [0..255];
    reverse_lens($list, $lens);
    $list->[0] * $list->[1];
}

sub solve_part_2($input) {
    knot_hash($input);
}

sub knot_hash($input) {
    my $lens = [ map ord, split //, $input ];
    push @$lens, 17, 31, 73, 47, 23;
    my $list = [0..255];
    my ($pos, $skip) = (0, 0);
    for (0..63) {
        ($pos, $skip) = reverse_lens($list, $lens, $pos, $skip);
    }
    my $dense_hash = dense_hash($list);
    my $hash;
    open my $sh, '>', \$hash;
    for my $d (@$dense_hash) {
        $sh->printf("%02x", $d);
    }
    close $sh;
    $hash;
}

sub dense_hash($list) {
    my @l;
    for my $i (0..15) {
        my $v = $list->[$i * 16];
        for my $k (1..15) {
            $v ^= $list->[$i * 16 + $k];
        }
        push @l, $v;
    }
    \@l;
}

sub reverse_lens($list, $lens, $pos //= 0, $skip //= 0) {
    for my $len (@$lens) {
        reverse_subarray($list, $pos, $len);
        $pos += $len + $skip;
        $skip++;
    }
    ($pos, $skip);
}


sub reverse_subarray($list, $pos, $len) {
    my ($s, $e) = ($pos, $pos + $len - 1);
    while ($s < $e) {
        my $si = $s % @$list;
        my $ei = $e % @$list;
        @$list[$si, $ei] = @$list[$ei, $si];
        $s++;
        $e--;
    }
}

sub get_input() {
    my $l = <DATA>;
    close DATA;
    chomp $l;
    $l;
}


__DATA__
165,1,255,31,87,52,24,113,0,91,148,254,158,2,73,153
