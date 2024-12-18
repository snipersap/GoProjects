# Message Formatter

As Textio evolves, the team has decided to introduce a new feature for custom message formats. Depending on the user's preferences, messages can be sent in different formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you'll implement a system using interfaces.
## Assignment

- Implement the Formatter interface with a method Format that returns a formatted string.
- Define structs that satisfy the Formatter interface: PlainText, Bold, Code.
    - The structs must all have a message field of type string

- PlainText should return the message as is.
- Bold should wrap the message in two asterisks (\**) to simulate bold text (e.g., **message** ).
- Code should wrap the message in a single backtick (\`) to simulate code block (e.g., `message`)

### Test results
```bash
=== RUN   TestSendMessage/Test_Case_1
---------------------------------
Test Case 1
Inputs:     ({Hello, World!})
Expecting:  Hello, World!
Actual:     Hello, World!
Pass
--- PASS: TestSendMessage/Test_Case_1 (0.00s)
=== RUN   TestSendMessage/Test_Case_2
---------------------------------
Test Case 2
Inputs:     ({Bold Message})
Expecting:  **Bold Message**
Actual:     **Bold Message**
Pass
--- PASS: TestSendMessage/Test_Case_2 (0.00s)
=== RUN   TestSendMessage/Test_Case_3
---------------------------------
Test Case 3
Inputs:     ({Code Message})
Expecting:  `Code Message`
Actual:     `Code Message`
Pass
--- PASS: TestSendMessage/Test_Case_3 (0.00s)
=== RUN   TestSendMessage/Test_Case_4
---------------------------------
Test Case 4
Inputs:     ({})
Expecting:  ``
Actual:     ``
Pass
--- PASS: TestSendMessage/Test_Case_4 (0.00s)
=== RUN   TestSendMessage/Test_Case_5
---------------------------------
Test Case 5
Inputs:     ({})
Expecting:  ****
Actual:     ****
Pass
--- PASS: TestSendMessage/Test_Case_5 (0.00s)
=== RUN   TestSendMessage/Test_Case_6
---------------------------------
Test Case 6
Inputs:     ({})
Expecting:
Actual:
Pass
--- PASS: TestSendMessage/Test_Case_6 (0.00s)
---------------------------------
6 passed, 0 failed
--- PASS: TestSendMessage (0.00s)
PASS
ok      message-formatter-textio        0.576s
```
