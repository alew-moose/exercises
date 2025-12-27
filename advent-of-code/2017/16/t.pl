#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;
use Benchmark qw(timethese);

# my $moves = get_input('test-input.txt');
my $moves = get_input('input.txt');
say 'part 1: ', solve_part_1($moves);
say 'part 2: ', solve_part_2($moves);

# timethese(1, {
#     v1 => sub { solve_part_2_v1($moves) },
#     v2 => sub { solve_part_2_v2($moves) },
#     v3 => sub { solve_part_2_v3($moves) },
#     v4 => sub { solve_part_2($moves) },
# });

sub solve_part_1($moves) {
    my $p = ['a'..'p'];
    for my $m (@$moves) {
        move($p, $m);
    }
    join '', @$p;
}

sub solve_part_2_v1($moves) {
    my $p = ['a'..'p'];
    for (1..1e3) {
        for my $m (@$moves) {
            move($p, $m);
        }
    }
    join '', @$p;
}

sub solve_part_2_v2($moves) {
    my $moves_opt = optimize_moves($moves);
    my $p = ['a'..'p'];
    for (1..1e3) {
        for my $m (@$moves_opt) {
            move($p, $m);
        }
    }
    join '', @$p;
}

sub solve_part_2_v3($moves) {
    my $p = [['a'..'p']];
    $p->[1] = { map { $_ => ord($_) - ord('a') } $p->[0]->@* };
    my $moves_opt = optimize_moves($moves);
    for (1..1e3) {
        for my $m (@$moves_opt) {
            move_v3($p, $m);
        }
    }
    join '', $p->[0]->@*;
}

# v4
sub solve_part_2($moves) {
    my $p = ['a'..'p'];
    my $orig = join '', @$p;
    my $n = 1;
    my $max = 1e9;
    for my $n (1..1000) {
        for my $m (@$moves) {
            move($p, $m);
        }
        my $s = join '', @$p;
        if ($s eq $orig) {
            $max = $max % $n;
            last;
        }
    }
    if ($max == 1e9) {
        die 'loop not found';
    }
    for (1..$max) {
        for my $m (@$moves) {
            move($p, $m);
        }
    }
    join '', @$p;
}

sub optimize_moves($moves) {
    my @moves_opt;
    my @moves_batch;
    for my $m (@$moves) {
        if ($m->[0] eq 'p') {
            if (@moves_batch) {
                push @moves_opt, optimize_moves_batch(@moves_batch);
                @moves_batch = ();
            }
            push @moves_opt, $m;
        } else {
            push @moves_batch, $m;
        }
    }
    if (@moves_batch) {
        push @moves_opt, optimize_moves_batch(@moves_batch);
    }
    \@moves_opt;
}

sub optimize_moves_batch(@moves) {
    return @moves unless grep { $_->[0] eq 's' } @moves;
    my $p = ['a'..'p'];
    for my $m (@moves) {
        move($p, $m);
    }
    my @idx;
    for my $c (@$p) {
        push @idx, ord($c) - ord('a');
    }
    ['i', \@idx];
}

sub move($p, $m) {
    if ($m->[0] eq 's') {
        spin($p, $m->[1]);
    } elsif ($m->[0] eq 'x') {
        exchange($p, $m->[1], $m->[2]);
    } elsif ($m->[0] eq 'p') {
        partner($p, $m->[1], $m->[2]);
    } elsif ($m->[0] eq 'i') {
        # "optimized"
        @$p = map { $p->[$_] } $m->[1]->@*;
    } else {
        die "unknown move $m->[0]";
    }
}

sub move_v3($p, $m) {
    if ($m->[0] eq 's') {
        unshift $p->[0]->@*, splice $p->[0]->@*, -$m->[1];
        rebuild_idx($p);
    } elsif ($m->[0] eq 'x') {
        my $array = $p->[0];
        my ($ai, $bi) = ($m->[1], $m->[2]);
        my $t = $array->[$ai];
        $array->[$ai] = $array->[$bi];
        $array->[$bi] = $t;
        $p->[1]{$array->[$ai]} = $ai;
        $p->[1]{$array->[$bi]} = $bi;
    } elsif ($m->[0] eq 'p') {
        my $array = $p->[0];
        my $hash = $p->[1];
        my ($a, $b) = ($m->[1], $m->[2]);
        my ($ai, $bi) = ($hash->{$a}, $hash->{$b});
        my $t = $array->[$ai];
        $array->[$ai] = $array->[$bi];
        $array->[$bi] = $t;
        $hash->{$a} = $bi;
        $hash->{$b} = $ai;
        # rebuild_idx($p);
    } elsif ($m->[0] eq 'i') {
        # "optimized"
        $p->[0]->@* = map { $p->[0][$_] } $m->[1]->@*;
        rebuild_idx($p);
    } else {
        die "unknown move $m->[0]";
    }
}

sub rebuild_idx($p) {
    while (my ($i, $c) = each $p->[0]->@*) {
        $p->[1]{$c} = $i;
    }
}

sub spin($p, $s) {
    unshift @$p, splice @$p, -$s;
}

sub exchange($p, $i, $k) {
    @$p[$i, $k] = @$p[$k, $i];
}

sub partner($p, $a, $b) {
    exchange($p, find($p, $a), find($p, $b));
}

sub find($p, $a) {
    for my $i (0..$#$p) {
        return $i if $p->[$i] eq $a;
    }
    die "'$a' not found";
}

sub get_input($file) {
    open my $fh, '<', $file or die "$file: $!";
    my $l = <$fh>;
    close $fh;
    [ map parse_cmd($_), split /,/, $l ];
}

sub parse_cmd($s) {
    $s =~ m{^s(\d+)|x(\d+)/(\d+)|p([a-p])/([a-p])}
        or die "failed to parse '$s'";
    return ['s', $1] if defined $1;
    return ['x', $2, $3] if defined $2;
    return ['p', $4, $5] if defined $4;
    die 'parse_cmd failed';
}

