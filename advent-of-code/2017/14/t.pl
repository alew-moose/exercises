#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;


# my $input = 'flqrgnkx'; # test input
my $input = 'ugkiagan';

say "part 1: ", solve_part_1($input);
say "part 2: ", solve_part_2($input);


sub solve_part_1($input) {
    my $used_cnt = 0;
    for my $row (0..127) {
        $used_cnt += pop_count(knot_hash("$input-$row"));
    }
    $used_cnt;
}

sub solve_part_2($input) {
    my $grid = make_grid($input);
    count_regions($grid);
}

sub count_regions($grid) {
    my $marked = [ map [ map 0, 0..127 ], 0..127 ];
    my $region = 1;
    for my $y (0..127) {
        for my $x (0..127) {
            if ($grid->[$y][$x] && !$marked->[$y][$x]) {
                mark_region($grid, $marked, $y, $x, $region);
                $region++;
            }
        }
    }
    $region - 1;
}

sub mark_region($grid, $marked, $y, $x, $region) {
    return if $y < 0 || $y > 127 || $x < 0 || $x > 127;
    return if $marked->[$y][$x] || !$grid->[$y][$x];
    $marked->[$y][$x] = $region;
    mark_region($grid, $marked, $y-1, $x, $region);
    mark_region($grid, $marked, $y+1, $x, $region);
    mark_region($grid, $marked, $y, $x-1, $region);
    mark_region($grid, $marked, $y, $x+1, $region);
}

sub make_grid($input) {
    my @rows;
    for my $row_n (0..127) {
        my $hash = knot_hash("$input-$row_n");
        my @row;
        for my $c (split //, $hash) {
            my $b = sprintf "%04b", hex $c;
            push @row, split //, $b;
        }
        push @rows, \@row;
    }
    \@rows;
}

sub pop_count($hash) {
    state %pc = do {
        my @chars = (0..9, 'a'..'f');
        my %cnt;
        while (my ($i, $c) = each @chars) {
            my $bin = sprintf "%b", $i;
            $cnt{$c} = $bin =~ tr/1//;
        }
        %cnt;
    };
    my $cnt = 0;
    for my $c (split //, $hash) {
        $cnt += $pc{$c};
    }
    $cnt;
}

# from day 10
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
