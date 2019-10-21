API response format:

1. General API response standard

```
----------------------------------
On Error:
----------------------------------
{
    'status': 'error',
    'code': '400|401|403|404|405|500|502|503|504',
    'message': []string,
    'data': [{}],
    'meta': {
	    'prev_url': nil,
        'next_url': nil,
        'data_count: 0,
    }
}

----------------------------------
On Success:
----------------------------------
{
    'status': 'success',
    'code': '200',
    'message': []string,
    'data': [
        {
            'question': 'string',
            'ans': string,
            'options': []string,
        }
    ],
    'meta': {
	    'prev_url': string,
        'next_url': string,
        'data_count: int,
    }
}
```