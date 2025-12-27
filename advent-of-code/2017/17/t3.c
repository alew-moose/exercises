#include <stdio.h>
#include <stdlib.h>

typedef struct list {
	int value;
	struct list *next;
} List;

typedef struct buff {
	size_t size;
	int idx;
	void *buff;
} Buff;

Buff new_buff(size_t size, int len)
{
	return (Buff){
		.size = size,
		.idx = 0,
		.buff = malloc(size * len),
	};
}

void *buff_alloc(Buff *b)
{
	return (char *)b->buff + b->size * b->idx++;
}

// ~3 min
int solve_part_2(int step)
{
	Buff b = new_buff(sizeof(List), 50000000);
	List *l = buff_alloc(&b);
	l->value = 0;
	l->next = l;
	int len = 1;
	for (int n = 0; n < 50000000; n++) {
		int val = l->value + 1;
		for (int k = 1; k <= step % len; k++) {
			l = l->next;
		}
		List *node = buff_alloc(&b);
		node->value = val;
		node->next = l->next;
		l->next = node;
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
