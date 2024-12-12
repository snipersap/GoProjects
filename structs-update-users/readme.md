# Update users

We need a way to differentiate between standard and premium users. When a new user is created, they need a membership type, and that type will determine the message character limit.
# Assignment

Create a new struct called Membership, it should have:

    A Type string field
    A MessageCharLimit integer field

Update the User struct to embed a Membership.

Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.

## Test Results
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


