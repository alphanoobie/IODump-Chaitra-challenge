#include <stdio.h>
int main()
{
    int number;
    int i;
    printf("Enter an number: ");
    scanf("%d", &number);  
    int sum = 0;
    for(i=1; i<=number; ++i){
        sum +=i;
    }
    printf("%d", sum);
    return 0;
}