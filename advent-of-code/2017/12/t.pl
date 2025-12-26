#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;


my $pipes = get_input('input.txt');

say 'part 1: ', solve_part_1($pipes);
say 'part 2: ', solve_part_2($pipes);

sub solve_part_1($pipes) {
    my $graph = make_graph($pipes);
    return group_size($graph, 0);
}

sub solve_part_2($pipes) {
    my $graph = make_graph($pipes);
    return groups_count($graph);
}


sub get_input($file) {
    my @pipes;
    open my $fh, '<', $file or die "$file: $!";
    while (<$fh>) {
        chomp;
        /(\d+)\s+<->\s+(.+)/ or die "failed to parse $_";
        my $from = $1;
        my @to = split /, /, $2, -1;
        push @pipes, [$from, \@to];
    }
    return \@pipes;
}

sub make_graph($pipes) {
    my %graph;
    for my $pipe (@$pipes) {
        my ($from, $tos) = @$pipe;
        for my $to (@$tos) {
            $graph{$from}{$to} = 1;
            $graph{$to}{$from} = 1;
        }
    }
    return \%graph;
}

sub group_size($graph, $node) {
    my $size = 0;
    my @queue = ($node);
    my %seen;
    while (@queue) {
        my $node = pop @queue;
        unless ($seen{$node}) {
            $size++;
            push @queue, keys $graph->{$node}->%*;
        }
        $seen{$node} = 1;
    }
    return $size;
}

sub groups_count($graph) {
    my %node_group;
    my $group = 1;
    for my $node (keys %$graph) {
        next if $node_group{$node};
        my @queue = ($node);
        while (@queue) {
            my $node = pop @queue;
            next if $node_group{$node};
            $node_group{$node} = $group;
            push @queue, keys $graph->{$node}->%*;
        }
        $group++;
    }
    return $group - 1;
}
