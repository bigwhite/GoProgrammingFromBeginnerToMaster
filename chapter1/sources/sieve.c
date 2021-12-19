#include <stdio.h>

#define LIMIT  50
#define PRIMES 10

void sieve() {
	int c, i,j,numbers[LIMIT], primes[PRIMES];

	for (i=0;i<LIMIT;i++){
		numbers[i]=i+2; /*fill the array with natural numbers*/
	}
	
	for (i=0;i<LIMIT;i++){
		if (numbers[i]!=-1){
			for (j=2*numbers[i]-2;j<LIMIT;j+=numbers[i])
				numbers[j]=-1; /*sieve the non-primes*/
		}
	}
	
	c = j = 0;
	for (i=0;i<LIMIT&&j<PRIMES;i++) {
		if (numbers[i]!=-1) {
			primes[j++] = numbers[i]; /*transfer the primes to their own array*/
			c++;
		}
	}
	
	for (i=0;i<c;i++) printf("%d\n",primes[i]);
}

int main() {
	sieve();
}
