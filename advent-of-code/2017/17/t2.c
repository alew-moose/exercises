#include <stdio.h>
#include <stdlib.h>

typedef struct list {
	int value;
	struct list *next;
} List;

void insert(List *l, int val)
{
	List *node = malloc(sizeof(List));
	node->value = val;
	node->next = l->next;
	l->next = node;
}

// ~2min 50sec
int solve_part_2(int step)
{
	List *l = malloc(sizeof(List));
	l->value = 0;
	l->next = l;
	int len = 1;
	for (int n = 0; n < 50000000; n++) {
		int val = l->value + 1;
		for (int k = 1; k <= step % len; k++) {
			l = l->next;
		}
		insert(l, val);
		len++;
		l = l->next;
	}
	while (l->value != 0) {
		l = l->next;
	}
	return l->next->value;
}

int main(void)
{
	int step = 377;
	printf("part 2: %d\n", solve_part_2(step));
}
