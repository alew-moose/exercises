#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

my ($max_depth, $scanners_range) = get_input('input.txt');

say "part 1: ", solve_part_1($max_depth, $scanners_range);
say "part 2: ", solve_part_2($max_depth, $scanners_range);

sub solve_part_1($max_depth, $scanners_range) {
    my %scanners_pos = map { $_ => 0 } keys %$scanners_range;
    my %scanners_dir = map { $_ => 1 } keys %$scanners_range;
    my $severity = 0;
    for my $depth (0..$max_depth) {
        if (exists($scanners_range->{$depth}) && $scanners_pos{$depth} == 0) {
            $severity += $depth * $scanners_range->{$depth};
        }
        for my ($depth, $range) (%$scanners_range) {
            $scanners_pos{$depth} += $scanners_dir{$depth};
            if ($scanners_pos{$depth} < 0) {
                $scanners_pos{$depth} = 1;
                $scanners_dir{$depth} = 1;
            } elsif ($scanners_pos{$depth} >= $scanners_range->{$depth}) {
                $scanners_pos{$depth} = $scanners_range->{$depth} - 2;
                $scanners_dir{$depth} = -1;
            }
        }
    }
    return $severity;
}

sub solve_part_2($max_depth, $scanners_range) {
    my $delay = 0;
    my @depth_period;
    for my ($depth, $range) (%$scanners_range) {
        push @depth_period, $depth, 2 * $range - 2;
    }
    DELAY: while (true) {
        for my ($depth, $period) (@depth_period) {
            if (($depth + $delay) % $period == 0) {
                $delay++;
                next DELAY;
            }
        }
        return $delay;
    }
}

sub get_input($file) {
    my %scanners_range;
    my $max_depth = 0;
    open my $fh, '<', $file or die "$file: $!";
    while (<$fh>) {
        /^(\d+): (\d+)$/ or die "failed to parse '$_'";
        my ($depth, $range) = ($1, $2);
        $max_depth = $depth if $depth > $max_depth;
        $scanners_range{$depth} = $range;
    }
    return $max_depth, \%scanners_range;
}
