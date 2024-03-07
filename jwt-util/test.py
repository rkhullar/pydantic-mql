'''
from jwt_util import ExtensionAdapter

adapter = ExtensionAdapter()
jwk = adapter.new_jwk()
print(jwk)
pem = adapter.jwk_to_pem(jwk)
print(pem)
'''

# '''
from pathlib import Path
import jwt


def read_data(name: str) -> str:
    path = Path(__file__).parent / 'data' / name
    with path.open('r') as f:
        return f.read().strip()


test_jwk = read_data('private-key.json')
test_pem = read_data('private-key.pem')

payload = {'message': 'hello world', 'count': 4}
result = jwt.encode(payload=payload, key=test_pem, algorithm='RS256', headers={'kid': 'asdf'})
print(result)
# '''