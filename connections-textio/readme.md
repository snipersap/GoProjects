# Connections

Textio has group chats that make communicating with multiple people much more efficient--if the chat doesn't descend into chaos. Instead of sending the message multiple times to individual people, you send one message to all of them at once.

## Assignment

Complete the `countConnections` function that takes an integer `groupSize` representing the number of people in the group chat and returns an integer representing the number of connections between them. For each additional person in the group, the number of new connections continues to grow. Use a `for` loop to accumulate the number of connections instead of directly using a mathematical formula.

### Output
```bash
Connections for 0 group size is:0
Connections for 1 group size is:0
Connections for 2 group size is:1
Connections for 3 group size is:3
Connections for 4 group size is:6
Connections for 5 group size is:10
Connections for 6 group size is:15
Connections for 7 group size is:21
Connections for 8 group size is:28
Connections for 9 group size is:36
Connections for 10 group size is:45
Connections for 11 group size is:55
Connections for 12 group size is:66
Connections for 13 group size is:78
Connections for 14 group size is:91
Connections for 15 group size is:105
Connections for 16 group size is:120
Connections for 17 group size is:136
Connections for 18 group size is:153
Connections for 19 group size is:171
```

### Test results
```bash
---------------------------------
Group Size: 1
Expecting: 0
Actual:    0
Pass
---------------------------------
Group Size: 2
Expecting: 1
Actual:    1
Pass
---------------------------------
Group Size: 3
Expecting: 3
Actual:    3
Pass
---------------------------------
Group Size: 4
Expecting: 6
Actual:    6
Pass
---------------------------------
Group Size: 0
Expecting: 0
Actual:    0
Pass
---------------------------------
Group Size: 10
Expecting: 45
Actual:    45
Pass
---------------------------------
Group Size: 100
Expecting: 4950
Actual:    4950
Pass
---------------------------------
Group Size: 1000
Expecting: 499500
Actual:    499500
Pass
---------------------------------
8 passed, 0 failed
PASS
ok      connections-textio      0.149s
```