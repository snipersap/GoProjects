# Update users

We need a way to differentiate between standard and premium users. When a new user is created, they need a membership type, and that type will determine the message character limit.
## Assignment

Create a new struct called Membership, it should have:

    A Type string field
    A MessageCharLimit integer field

Update the User struct to embed a Membership.

Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.

## Test results
```bash
$> go run main.go
User: johhny , Limit: 1000
User: Silke , Limit: 100

$> go test
---------------------------------
Test Passed:
Inputs:     (name: Syl, membershipType: standard)
Expecting:  Syl, standard, 100
Actual:     Syl, standard, 100
---------------------------------
Test Passed:
Inputs:     (name: Pattern, membershipType: premium)
Expecting:  Pattern, premium, 1000
Actual:     Pattern, premium, 1000
---------------------------------
Test Passed:
Inputs:     (name: Pattern, membershipType: standard)
Expecting:  Pattern, standard, 100
Actual:     Pattern, standard, 100
---------------------------------
Test Passed:
Inputs:     (name: Renarin, membershipType: standard)
Expecting:  Renarin, standard, 100
Actual:     Renarin, standard, 100
---------------------------------
Test Passed:
Inputs:     (name: Lift, membershipType: premium)
Expecting:  Lift, premium, 1000
Actual:     Lift, premium, 1000
---------------------------------
Test Passed:
Inputs:     (name: Dalinar, membershipType: standard)
Expecting:  Dalinar, standard, 100
Test Passed:
Inputs:     (name: Dalinar, membershipType: standard)
Test Passed:
Test Passed:
Test Passed:
Inputs:     (name: Dalinar, membershipType: standard)
Expecting:  Dalinar, standard, 100
Actual:     Dalinar, standard, 100
---------------------------------
6 passed, 0 failed
PASS
ok      structs-update-users    0.435s
```

# Send Message
## Assignment

Create a SendMessage method for the User struct.

It should take a message string and messageLength int as inputs.

If the messageLength is within the user's message character limit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false.

## Test results
```bash
$> go run .\main.go
User: johhny , Limit: 1000
User: Silke , Limit: 100
Message sent from Jhonny: Hello, I am Jhonny
Message too large to send for  Silke

$> go test
---------------------------------
Test Passed:
* user:               Syl
* membership type:    standard
* message:            Hello, Kaladin!
* expected result:    Hello, Kaladin!
* expected success:   true
* actual result:      Hello, Kaladin!
* actual success:     true
---------------------------------
Test Passed:
* user:               Pattern
* membership type:    premium
* message:            You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.
* expected result:    You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.
* expected success:   true
* actual result:      You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.
* actual success:     true
---------------------------------
Test Passed:
* user:               Dalinar
* membership type:    standard
* message:            I will take responsibility for what I have done. If I must fall, I will rise each time a better man.
* expected result:    I will take responsibility for what I have done. If I must fall, I will rise each time a better man.
* expected success:   true
* actual result:      I will take responsibility for what I have done. If I must fall, I will rise each time a better man.
* actual success:     true
---------------------------------
Test Passed:
* user:               Pattern
* membership type:    standard
* message:            Humans can see the world as it is not. It is why your lies can be so strong. You are able to not admit that they are lies.
* expected result:
* expected success:   false
* actual result:
* actual success:     false
---------------------------------
Test Passed:
* user:               Dabbid
* membership type:    premium
* message:            .........................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................
* expected result:
* expected success:   false
* actual result:
* actual success:     false
---------------------------------
5 passed, 0 failed
---------------------------------
Test Passed:
Inputs:     (name: Syl, membershipType: standard)
Expecting:  Syl, standard, 100
Actual:     Syl, standard, 100
---------------------------------
Test Passed:
Inputs:     (name: Pattern, membershipType: premium)
Expecting:  Pattern, premium, 1000
Actual:     Pattern, premium, 1000
---------------------------------
Test Passed:
Inputs:     (name: Pattern, membershipType: standard)
Expecting:  Pattern, standard, 100
Actual:     Pattern, standard, 100
---------------------------------
Test Passed:
Inputs:     (name: Renarin, membershipType: standard)
Expecting:  Renarin, standard, 100
Actual:     Renarin, standard, 100
---------------------------------
Test Passed:
Inputs:     (name: Lift, membershipType: premium)
Expecting:  Lift, premium, 1000
Actual:     Lift, premium, 1000
---------------------------------
Test Passed:
Inputs:     (name: Dalinar, membershipType: standard)
Expecting:  Dalinar, standard, 100
Actual:     Dalinar, standard, 100
---------------------------------
6 passed, 0 failed
PASS
ok      structs-update-users    0.280s
```
