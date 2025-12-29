#!/usr/bin/env perl
use v5.42;
use utf8;
use List::Util qw(sum);
use Carp qw(croak);
use DDP caller_info => 1;

if (($ARGV[0] // '') eq '-t') {
    test();
    exit;
}
if (($ARGV[0] // '') eq '-p') {
    find_permutation_methods();
    exit;
}

# my $rules = get_input('test-input.txt');
my $rules = get_input('input.txt');
my $transformer = make_transformer($rules);

say 'part 1: ', solve($transformer, 5);
say 'part 2: ', solve($transformer, 18);

sub solve($transformer, $iterations) {
    my $block = line_to_block('.#./..#/###');
    for (1..$iterations) {
        $block = transform_block($block, $transformer);
    }
    block_count_elems($block, '#');
}

sub transform_block($block, $transformer) {
    my $subblock_size;
    if (@$block % 2 == 0) {
        $subblock_size = 2;
    } elsif (@$block % 3 == 0) {
        $subblock_size = 3;
    }
    else {
        die "block is not divisible by 2 or 3";
    }
    my $new_subblock_size = $subblock_size + 1;
    my $new_block = new_block(@$block / $subblock_size * $new_subblock_size);
    block_iterate_subblocks($block, $subblock_size, sub($by, $bx, $subblock) {
        my $new_subblock = $transformer->(block_to_line($subblock));
        my $subblock_line = block_to_line($new_subblock);
        block_set_subblock($new_block, $by, $bx, $new_subblock);
    });
    $new_block;
}

sub block_iterate_subblocks($block, $subblock_size, $cb) {
    my $subblock = new_block($subblock_size);
    my $size = @$block / $subblock_size;
    for my $by (0..$size-1) {
        for my $bx (0..$size-1) {
            for my $y (0..$subblock_size-1) {
                for my $x (0..$subblock_size-1) {
                    $subblock->[$y][$x] = $block->[$by*$subblock_size+$y][$bx*$subblock_size+$x];
                }
            }
            $cb->($by, $bx, $subblock);
        }
    }
}

sub block_set_subblock($block, $by, $bx, $subblock) {
    for my $sby (0..$#$subblock) {
        my $y = $by * @$subblock + $sby;
        for my $sbx (0..$#$subblock) {
            my $x = $bx * @$subblock + $sbx;
            $block->[$y][$x] = $subblock->[$sby][$sbx];
        }
    }
}

sub block_count_elems($block, $elem) {
    sum(map { scalar grep { $_ eq $elem } @$_ } @$block);
}

sub make_transformer($rules) {
    my %line_to_block;
    for my $rule (@$rules) {
        my ($from_line, $to_line) = @$rule;
        my $from_block = line_to_block($from_line);
        my $to_block = line_to_block($to_line);
        for my $from_block_perm (@{block_permutations($from_block)}) {
            my $from_line_perm = block_to_line($from_block_perm);
            if ($line_to_block{$from_line_perm}) {
                my $b = $line_to_block{$from_line_perm};
                if (block_to_line($b) ne $to_line) {
                    die "conflicting rule for $from_line_perm";
                }
            }
            $line_to_block{$from_line_perm} = $to_block;
        }
    }
    sub($line) {
        $line_to_block{$line} // die "no block for '$line'";
    };
}

sub block_permutations($block) {
    no strict 'refs';
    state @methods = (
        [qw(flip_h)],
        [qw(flip_h flip_v)],
        [qw(flip_h flip_v rot_r)],
        [qw(flip_h rot_r)],
        [qw(flip_v)],
        [qw(flip_v rot_r)],
        [qw(rot_r)],
    );
    my @perms = ($block);
    for my $method_names (@methods) {
        my $block_copy = block_copy($block);
        for my $method_name (@$method_names) {
            my $method_sub = \&{"block_$method_name"};
            $block_copy = $method_sub->($block_copy);
        }
        push @perms, $block_copy;
    }
    \@perms;
}

sub block_flip_h($block) {
    my $b = block_copy($block);
    my ($s, $e) = (0, $#$b);
    while ($s < $e) {
        @$b[$s, $e] = @$b[$e, $s];
        $s++;
        $e--;
    }
    $b;
}

sub block_flip_v($block) {
    my $b = block_copy($block);
    for my $y (0..$#$b) {
        my ($xs, $xe) = (0, $#{$b->[0]});
        while ($xs < $xe) {
            @{$b->[$y]}[$xs, $xe] = @{$b->[$y]}[$xe, $xs];
            $xs++;
            $xe--;
        }
    }
    $b;
}

sub block_rot_r($block) {
    my $size = @$block;
    my $b = new_block($size);
    for my $y (0..$size-1) {
        for my $x (0..$size-1) {
            $b->[$y][$x] = $block->[$size-$x-1][$y];
        }
    }
    $b;
}

sub find_permutation_methods {
    # my $line = '12/34';
    my $line = '123/456/789';
    my $block = line_to_block($line);
    my %lines_methods = (
        $line => ['orig'],
    );
    find_perms(\%lines_methods, [], $block);

    say '';
    my @methods_strs;
    for my ($line) (sort keys %lines_methods) {
        my $methods = $lines_methods{$line};
        my $methods_str = join('+', @$methods);
        push @methods_strs, $methods_str;
        say $methods_str;
        say block_to_str(line_to_block($line));
        say '';
    }

    say "methods:";
    say for sort @methods_strs;
}

sub find_perms($lines_methods, $methods, $block) {
    no strict 'refs';
    for my $method_name (qw(flip_h flip_v rot_r)) {
        if (@$methods && $methods->[-1] =~ /flip/ && $methods->[-1] eq $method_name) {
            next;
        }
        my $method_sub = \&{"block_$method_name"};
        my $new_block = $method_sub->($block);
        my $new_line = block_to_line($new_block);
        my $new_methods = [@$methods, $method_name];
        if (my $prev_methods = $lines_methods->{$new_line}) {
            printf(
                "%s == %s (%s)\n",
                join('+', @$new_methods),
                join('+', @$prev_methods),
                $new_line,
            );
            if (@$new_methods >= @$prev_methods) {
                next;
                $lines_methods->{$new_line} = [@$new_methods];
            }
        }
        $lines_methods->{$new_line} = [@$new_methods];
        find_perms($lines_methods, $new_methods, $new_block);
    }
}

sub line_to_block($line) {
    my $block = [ map { [ split //, $_ ] } split /\//, $line ];
    for my $row (@$block) {
        croak "$line is not square" if @$row != @$block;
    }
    $block;
}

sub block_to_line($block) {
    join '/', map { join '', @$_ } @$block;
}

sub block_to_str($block) {
    join "\n", map { join '', @$_ } @$block;
}

sub new_block($size) {
    [
        map { [ map { undef } 1..$size ] }
        1..$size
    ];
}

sub block_copy($block) {
    [ map { [ @$_ ] } @$block ];
}

sub get_input($file) {
    state $s = qr/[.#]/;
    state $rule_2_to_3 = qr{($s{2}/$s{2}) => ($s{3}/$s{3}/$s{3})};
    state $rule_3_to_4 = qr{($s{3}/$s{3}/$s{3}) => ($s{4}/$s{4}/$s{4}/$s{4})};
    state $rule_re = qr/^(?|$rule_2_to_3|$rule_3_to_4)$/;
    my @rules;
    open my $fh, '<', $file or die "$file: $!";
    while (<$fh>) {
        chomp;
        /$rule_re/ or die "failed to parse $_";
        push @rules, [$1, $2];
    }
    close $fh;
    \@rules;
}

sub test {
    use Test::More;

    {
        my $block = line_to_block('12/34');
        is_deeply $block, [[1, 2], [3, 4]];
        is block_to_line($block), '12/34';

        my $new_block = block_rot_r($block);
        is_deeply $new_block, line_to_block('31/42');
        is_deeply $block, line_to_block('12/34');
        $new_block = block_rot_r($new_block);
        is block_to_line($new_block), '43/21';
        $new_block = block_rot_r($new_block);
        is block_to_line($new_block), '24/13';
        $new_block = block_rot_r($new_block);
        is block_to_line($new_block), '12/34';

        $new_block = block_flip_v($block);
        is block_to_line($new_block), '21/43';
        is block_to_line($block), '12/34';

        $new_block = block_flip_v($new_block);
        is block_to_line($new_block), '12/34';

        $new_block = block_flip_h($block);
        is block_to_line($new_block), '34/12';
        is block_to_line($block), '12/34';
        $new_block = block_flip_h($new_block);
        is block_to_line($new_block), '12/34';
    }

    {
        my $block = line_to_block('123/456/789');
        is_deeply $block, [[1, 2, 3], [4, 5, 6], [7, 8, 9]];
        is_deeply block_copy($block), $block;

        my $new_block = block_rot_r($block);
        is block_to_line($new_block), '741/852/963';
        is block_to_line($block), '123/456/789';

        is block_to_line(block_flip_h($new_block)), '963/852/741';

        is block_to_line(block_flip_v($new_block)), '147/258/369';
    }

    {
        no strict 'refs';
        my @tests = (
            [[qw(flip_h rot_r)], '13/24'],
            [[qw(flip_v)], '21/43'],
            [[qw(flip_h flip_v rot_r)], '24/13'],
            [[qw(rot_r)], '31/42'],
            [[qw(flip_h)], '34/12'],
            [[qw(flip_v rot_r)], '42/31'],
            [[qw(flip_h flip_v)], '43/21'],
            [[qw(rot_r flip_v)], '13/24'],
            [[qw(rot_r rot_r rot_r)], '24/13'],
            [[qw(rot_r flip_h)], '42/31'],
            [[qw(rot_r rot_r)], '43/21'],
        );
        for my $test (@tests) {
            my ($method_names, $expected_line) = @$test;
            my $block = line_to_block('12/34');
            for my $method_name (@$method_names) {
                my $method_sub = \&{"block_$method_name"};
                $block = $method_sub->($block);
            }
            local $" = '+';
            is block_to_line($block), $expected_line, "@$method_names -> $expected_line";
        }
    }

    {
        no strict 'refs';
        my @tests = (
            [[qw(flip_h rot_r)], '147/258/369'],
            [[qw(flip_v)], '321/654/987'],
            [[qw(flip_h flip_v rot_r)], '369/258/147'],
            [[qw(rot_r)], '741/852/963'],
            [[qw(flip_h)], '789/456/123'],
            [[qw(flip_v rot_r)], '963/852/741'],
            [[qw(flip_h flip_v)], '987/654/321'],
            [[qw(rot_r flip_v)], '147/258/369'],
            [[qw(rot_r rot_r rot_r)], '369/258/147'],
            [[qw(rot_r flip_h)], '963/852/741'],
            [[qw(rot_r rot_r)], '987/654/321'],
        );
        for my $test (@tests) {
            my ($method_names, $expected_line) = @$test;
            my $block = line_to_block('123/456/789');
            for my $method_name (@$method_names) {
                my $method_sub = \&{"block_$method_name"};
                $block = $method_sub->($block);
            }
            local $" = '+';
            is block_to_line($block), $expected_line, "@$method_names -> $expected_line";
        }
    }

    done_testing();
}
