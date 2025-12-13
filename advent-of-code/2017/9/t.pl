#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;

# my $input_str = read_input('test-input.txt');
my $input_str = read_input('input.txt');

say "part 1: ", solve_part_1($input_str);
say "part 2: ", solve_part_2($input_str);

sub read_input($file) {
    open my $fh, '<:encoding(UTF-8)', $file or die "$file: $!";
    my $line = <$fh>;
    chomp $line;
    $line;
}

sub solve_part_1($input_str) {
    my $score = 0;
    parse(
        input           => $input_str,
        group_enter_cb  => sub ($level) { $score += $level },
        garbage_char_cb => sub ($char) {},
    );
    $score;
}

sub solve_part_2($input_str) {
    my $garbage_chars_cnt = 0;
    parse(
        input           => $input_str,
        group_enter_cb  => sub ($level) {},
        garbage_char_cb => sub ($char) { $garbage_chars_cnt++ },
    );
    $garbage_chars_cnt;
}

use constant {
    OUTSIDE        => 0,
    INSIDE_GROUP   => 1,
    INSIDE_GARBAGE => 2,
};

sub parse(%param) {
    my @chars = split //, $param{input};
    my $state = OUTSIDE;
    my $level = 0;
    my $skip = false;
    while(my ($pos, $c) = each @chars) {
        if ($skip) {
            $skip = false;
            next;
        }
        if ($state == OUTSIDE) {
            if ($c eq '{') {
                $state = INSIDE_GROUP;
                $level++;
                $param{group_enter_cb}->($level);
            } elsif ($c eq '<') {
                $state = INSIDE_GARBAGE;
            } elsif ($c eq ',') {
                # no op
            } else {
                parse_fail($input_str, $pos, $state);
                STDERR->say($input_str);
                STDERR->print(' ' x $pos);
                STDERR->say("^");
                die "unexpected char '$c' at position $pos in state $state";
            }
        } elsif ($state == INSIDE_GROUP) {
            if ($c eq '{') {
                $level++;
                $param{group_enter_cb}->($level);
            } elsif ($c eq '<') {
                $state = INSIDE_GARBAGE;
            } elsif ($c eq '}') {
                $level--;
                if ($level == 0) {
                    $state = OUTSIDE;
                } elsif ($level < 0) {
                    die "level < 0 at position $pos";
                }
            }
        } elsif ($state == INSIDE_GARBAGE) {
            if ($c eq '>') {
                $state = $level > 0 ? INSIDE_GROUP : OUTSIDE;
            } elsif ($c eq '!') {
                $skip = true;
            } else {
                $param{garbage_char_cb}->($c);
            }
        } else {
            die "invalid state $state at position $pos";
        }
    }
    die "unexpected state $state at the end of input" if $state != OUTSIDE;
}
