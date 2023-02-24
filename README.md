# lego_sawmills


Run by 

```
go run main/main.go
```
and enter the input as specified in the Task:

1. The first line of input for each test case contains a single integer Z, the number of sawmills (each is
connected to its own river) in the test case.
2. This is followed by Z lines, each describing the tree trunks.
3. The first number in each line is the number E of tree trunks that are cut.
4. Following it are E strict positive integers, indicating the length (in blocks) of the tree trunks.
5. The input is terminated by a description starting with Z = 0.


### Assumptions:

I think the task description is not entirely clear about how the sawmills work. 
I assume that it will only sell whenever the sawnmill makes a cut.
For example if we have:
1 2 1 then it will be 3 (first is remainder and then it makes the only cut at 2) 

And that we can sell more than once per tree trunk. 
For example:
Or if we have 5 4 3 then it will be 1 + (-1 + 1) + 1 = 2


