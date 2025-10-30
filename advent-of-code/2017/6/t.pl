#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

my $input = '14	0	15	12	11	11	3	5	1	6	8	4	9	1	8	4';
my $blocks = [split ' ', $input, -1];

say 'part 1: ', solve_part_1($blocks);
say 'part 2: ', solve_part_2($blocks);

sub solve_part_1($blocks_ref) {
    my $blocks = [@$blocks_ref];
    my %seen;
    $seen{join '-', @$blocks}++;
    my $n = 0;
    while (true) {
        redistribute($blocks);
        $n++;
        last if $seen{join '-', @$blocks}++;
    }
    $n;
}

sub solve_part_2($blocks_ref) {
    my $blocks = [@$blocks_ref];
    my %seen;
    $seen{join '-', @$blocks}++;
    my $n = -1;
    while (true) {
        redistribute($blocks);
        my $key = join '-', @$blocks;
        $seen{$key}++;
        if ($n == -1 && $seen{$key} == 2) {
            $n = 0;
        } elsif ($n >= 0) {
            $n++;
        }
        if ($seen{$key} == 3) {
            return $n;
        }
    }
    $n;
}

sub redistribute($blocks) {
    my $max_i = max_idx($blocks);
    my $n = $blocks->[$max_i];
    $blocks->[$max_i] = 0;

    my $nb = int($n / @$blocks);
    $_ += $nb for @$blocks;
    $n = $n % @$blocks;

    my $i = ($max_i + 1) % @$blocks;
    while ($n > 0) {
        $blocks->[$i]++;
        $n--;
        $i = ($i + 1) % @$blocks;
    }
}

sub max_idx($nums) {
    my $max_i = 0;
    for my $i (1..$#$nums) {
        $max_i = $i if $nums->[$i] > $nums->[$max_i];
    }
    $max_i;
}
