# Domain Error Dictionary

This document provides a comprehensive reference for all domain-level errors in the proletariat-budget-core application. These errors are defined across various domain entities and represent business rule violations or invalid states.

## Error Reference Table

| Error Name | Returned Text | Description | Possible Cases |
|------------|---------------|-------------|----------------|
| **Account Errors** |
| `ErrAccountNotFound` | "account not found" | The requested account does not exist in the system | Account lookup by ID fails, account was deleted, invalid account reference |
| `ErrAccountInactive` | "account is inactive" | Operation attempted on an inactive account | Trying to perform transactions on deactivated accounts, accessing disabled account features |
| `ErrInsufficientBalance` | "insufficient account balance" | Account balance is too low for the requested operation | Withdrawal exceeds available balance, transfer amount greater than account funds |
| `ErrInvalidAccountType` | "invalid account type" | The specified account type is not recognized or allowed | Creating account with unsupported type, invalid account type in business logic |
| `ErrAccountHasTransactions` | "cannot delete account with existing transactions" | Account deletion blocked due to existing transaction history | Attempting to delete account with transaction records, data integrity protection |
| `ErrAccountAlreadyActive` | "account is already active" | Attempting to activate an already active account | Redundant activation operations, state validation failures |
| `ErrAccountAlreadyInactive` | "account is already inactive" | Attempting to deactivate an already inactive account | Redundant deactivation operations, state validation failures |
| `ErrInvalidCurrency` | "invalid currency" | The specified currency is not supported or recognized | Using unsupported currency codes, invalid currency format |
| `ErrAccountHasActiveRecurrencePatterns` | "account has active recurrence patterns" | Account cannot be modified/deleted due to active recurring transactions | Deleting account with scheduled recurring payments, modifying account with active patterns |
| `ErrAccountHasActiveSavingsGoals` | "account has active savings goals" | Account cannot be modified/deleted due to linked active savings goals | Deleting account with ongoing savings goals, account closure with active goals |
| **Transaction Errors** |
| `ErrTransactionNotFound` | "transaction not found" | The requested transaction does not exist | Transaction lookup by ID fails, transaction was deleted, invalid transaction reference |
| `ErrTransactionAmountMustBePositive` | "transaction amount must be positive" | Transaction amount validation failed | Creating transaction with negative or zero amount, invalid amount input |
| `ErrTransactionRollbackAmountMustBeNegative` | "rollback amount must be negative" | Rollback transaction requires negative amount | Incorrect rollback amount format, validation of reversal transactions |
| `ErrTransactionDescriptionEmpty` | "description cannot be empty" | Transaction description is required but not provided | Creating transaction without description, empty description validation |
| `ErrTransactionCurrencyEmpty` | "currency cannot be empty" | Transaction currency is required but not provided | Creating transaction without currency, missing currency validation |
| **Category Errors** |
| `ErrCategoryNotFound` | "category not found" | The requested category does not exist in the system | Category lookup by ID fails, category was deleted, invalid category reference |
| `ErrCategoryHasActiveSavingsGoals` | "category has active savings goals and cannot be deleted" | Category deletion blocked due to linked active savings goals | Deleting category with ongoing savings goals, data integrity protection |
| `ErrCategoryInactive` | "category is inactive" | Operation attempted on an inactive category | Using deactivated category for transactions, accessing disabled category |
| `ErrCategoryAlreadyActive` | "category is already active" | Attempting to activate an already active category | Redundant activation operations, state validation failures |
| `ErrCategoryAlreadyInactive` | "category is already inactive" | Attempting to deactivate an already inactive category | Redundant deactivation operations, state validation failures |
| `ErrCategoryUsedInExpenditure` | "category is used in expenditures" | Category cannot be deleted due to existing expenditure records | Deleting category with transaction history, data integrity protection |
| `ErrCategoryUsedInTransfer` | "category is used in transfers" | Category cannot be deleted due to existing transfer records | Deleting category with transfer history, data integrity protection |
| `ErrCategoryUsedInSavingGoal` | "category is used in saving goals" | Category cannot be deleted due to existing savings goal associations | Deleting category linked to savings goals, data integrity protection |
| `ErrCategoryUsedInIngress` | "category is used in ingresses" | Category cannot be deleted due to existing income records | Deleting category with income transaction history, data integrity protection |
| `ErrCategoryUsedInEntity` | "category is used in entity" | Category cannot be deleted due to existing entity associations | Deleting category with general entity relationships, data integrity protection |
| **Tag Errors** |
| `ErrTagNotFound` | "tag not found" | The requested tag does not exist in the system | Tag lookup by ID fails, tag was deleted, invalid tag reference |
| `ErrTagInUse` | "tag is in use and cannot be deleted" | Tag deletion blocked due to existing associations | Deleting tag with active references, data integrity protection |
| `ErrUnknownTagType` | "unknown tag type" | The specified tag type is not recognized | Using unsupported tag type, invalid tag type validation |
| `ErrTagAlreadyExists` | "tag with this name and type already exists" | Duplicate tag creation attempt | Creating tag with existing name-type combination, uniqueness constraint violation |
| `ErrTagNameEmpty` | "tag name cannot be empty" | Tag name is required but not provided | Creating tag without name, empty name validation |
| **Exchange Rate Errors** |
| `ErrExchangeRateNotFound` | "exchange rate not found" | The requested exchange rate does not exist | Exchange rate lookup fails, rate was deleted, invalid rate reference |
| `ErrExchangeRateSellRateMustBePositive` | "sell rate must be positive" | Sell rate validation failed | Creating exchange rate with negative or zero sell rate, invalid rate input |
| `ErrExchangeRateBuyRateMustBePositive` | "buy rate must be positive" | Buy rate validation failed | Creating exchange rate with negative or zero buy rate, invalid rate input |
| `ErrExchangeRateFromCurrencyNil` | "from currency cannot be nil" | Source currency is required but not provided | Creating exchange rate without source currency, missing currency validation |
| `ErrExchangeRateToCurrencyNil` | "to currency cannot be nil" | Target currency is required but not provided | Creating exchange rate without target currency, missing currency validation |
| `ErrExchangeRateSameCurrency` | "from and to currencies cannot be the same" | Exchange rate cannot be created between identical currencies | Creating rate with same source and target currency, business logic violation |
| **Expenditure Errors** |
| `ErrExpenditureNotFound` | "expenditure not found" | The requested expenditure does not exist | Expenditure lookup by ID fails, expenditure was deleted, invalid expenditure reference |
| `ErrExpenditureCategoryNil` | "category is required" | Expenditure category is required but not provided | Creating expenditure without category, missing category validation |
| `ErrExpenditureTransactionNil` | "transaction is required" | Expenditure transaction is required but not provided | Creating expenditure without transaction, missing transaction validation |
| **Household Member Errors** |
| `ErrMemberHasActiveAccounts` | "member has active accounts" | Member cannot be deleted due to existing active accounts | Deleting member with active account associations, data integrity protection |
| `ErrMemberAlreadyActive` | "member is already active" | Attempting to activate an already active member | Redundant activation operations, state validation failures |
| `ErrMemberNotFound` | "member not found" | The requested household member does not exist | Member lookup by ID fails, member was deleted, invalid member reference |
| `ErrMemberAlreadyInactive` | "member is already inactive" | Attempting to deactivate an already inactive member | Redundant deactivation operations, state validation failures |
| `ErrMemberInactive` | "member is inactive" | Operation attempted on an inactive member | Using deactivated member for operations, accessing disabled member |
| **Authentication Errors** |
| `ErrUserNotFound` | "user not found" | The requested user does not exist in the system | User lookup fails, user was deleted, invalid user reference |
| `ErrInvalidCredentials` | "invalid credentials" | Authentication failed due to incorrect credentials | Wrong password, invalid username, authentication failure |
| `ErrUserAlreadyExists` | "user already exists" | User creation failed due to existing user | Duplicate user registration, email/username already taken |
| `ErrUserInactive` | "user is inactive" | Operation attempted on an inactive user account | Accessing disabled user account, deactivated user operations |
| `ErrInvalidToken` | "invalid token" | Authentication token is invalid or expired | Expired JWT token, malformed token, token validation failure |
| `ErrTokenExpired` | "token has expired" | Authentication token has exceeded its validity period | Session timeout, expired access token, token refresh needed |
| `ErrInsufficientPermissions` | "insufficient permissions" | User lacks required permissions for the operation | Role-based access control violation, unauthorized operation attempt |
| `ErrInvalidRole` | "invalid role" | The specified user role is not recognized | Assigning unsupported role, invalid role validation |