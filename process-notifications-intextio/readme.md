# Process Notifications

Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.
## Assignment

Implement the importance methods for each message type. They should return the importance score for each message type.

- For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent, the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
- For a groupMessage the importance score is based on the messages priorityLevel
- All systemAlert messages should return a 100 importance score.

Complete the processNotification function. It should identify the type and return different values for each type

- For a directMessage, return the sender's username and importance score.
- For a groupMessage, return the group's name and the importance score.
- For an systemAlert, return the alert code and the importance score.
- If the notification does not match any known type, return an empty string and a score of 0.

### Test results
```bash
=== RUN   Test/TestProcessNotification_1
---------------------------------
Test Passed: TestProcessNotification_1
Notification: {senderUsername:Kaladin messageContent:Life before death priorityLevel:10 isUrgent:true}
Expecting:    Kaladin/50
Actual:       Kaladin/50
Pass
--- PASS: Test/TestProcessNotification_1 (0.00s)
=== RUN   Test/TestProcessNotification_2
---------------------------------
Test Passed: TestProcessNotification_2
Notification: {groupName:Bridge 4 messageContent:Soups ready! priorityLevel:2}
Expecting:    Bridge 4/2
Actual:       Bridge 4/2
Pass
--- PASS: Test/TestProcessNotification_2 (0.00s)
=== RUN   Test/TestProcessNotification_3
---------------------------------
Test Passed: TestProcessNotification_3
Notification: {alertCode:ALERT001 messageContent:THIS IS NOT A TEST HIGH STORM COMING SOON}
Expecting:    ALERT001/100
Actual:       ALERT001/100
Pass
--- PASS: Test/TestProcessNotification_3 (0.00s)
=== RUN   Test/TestProcessNotification_4
---------------------------------
Test Passed: TestProcessNotification_4
Notification: {senderUsername:Shallan messageContent:I am that I am. priorityLevel:5 isUrgent:false}
Expecting:    Shallan/5
Actual:       Shallan/5
Pass
--- PASS: Test/TestProcessNotification_4 (0.00s)
=== RUN   Test/TestProcessNotification_5
---------------------------------
Test Passed: TestProcessNotification_5
Notification: {groupName:Knights Radiant messageContent:For the greater good. priorityLevel:10}
Expecting:    Knights Radiant/10
Actual:       Knights Radiant/10
Pass
--- PASS: Test/TestProcessNotification_5 (0.00s)
=== RUN   Test/TestProcessNotification_6
---------------------------------
Test Passed: TestProcessNotification_6
Notification: {senderUsername:Adolin messageContent:Duels are my favorite. priorityLevel:3 isUrgent:true}
Expecting:    Adolin/50
Actual:       Adolin/50
Pass
--- PASS: Test/TestProcessNotification_6 (0.00s)
---------------------------------
6 passed, 0 failed
--- PASS: Test (0.00s)
PASS
ok      process-notifications-intextio  (cached)
```

