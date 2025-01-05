# User Input

In Textio, users can set their profile status to communicate their current activity to those that choose to read it... However, there are some restrictions on what these statuses can contain. Your task is to implement a function that validates a user's status update. The status update cannot be empty and must not exceed 140 characters.
## Assignment

Complete the validateStatus function. It should return an error when the status update violates any of the rules:

- If the status is empty, return an error that reads status cannot be empty.
- If the status exceeds 140 characters, return an error that says status exceeds 140 characters.

### Test result
```bash
=== RUN   TestValidateStatus
---------------------------------
Inputs:     ""
Expecting:  "status cannot be empty"
Actual:     "status cannot be empty"
Pass
---------------------------------
Inputs:     "This is a valid status update that is well within the character limit."
Expecting:  ""
Actual:     ""
Pass
---------------------------------
Inputs:     "This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco."
Expecting:  "status exceeds 140 characters"
Actual:     "status exceeds 140 characters"
Pass
---------------------------------
Inputs:     "Another valid status."
Expecting:  ""
Actual:     ""
Pass
---------------------------------
Inputs:     "This status update, while derivative, contains exactly one hundred and forty-one characters, which is over the status update character limit."
Expecting:  "status exceeds 140 characters"
Actual:     "status exceeds 140 characters"
Pass
---------------------------------
5 passed, 0 failed
--- PASS: TestValidateStatus (0.00s)
PASS
ok      validatestatus-Textio   (cached)
```