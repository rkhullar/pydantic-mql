from pydantic_mql import Condition

check_rating = {'operator': '$lte', 'field': 'rating', 'value': 80}
check_label = {'operator': '$eq', 'field': 'label', 'value': 'lab'}
condition_data = {'operator': '$and', 'conditions': [check_rating, check_label]}

test_data = check_label

condition = Condition.model_validate(test_data)
print(condition)
print(condition.model_dump_json(indent=4))

test_data = {'rating': 90, 'label': 'lab'}
print(condition.eval(test_data))

