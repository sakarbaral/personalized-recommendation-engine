import pandas as pd
import numpy as np
import time


num_records = 100

data = {
    'user_id': [f'user_{i}' for i in range(num_records)],
    'event_type': np.random.choice(['view', 'click', 'purchase'], size=num_records),
    'timestamp': [int(time.time()) - np.random.randint(0, 10000000) for _ in range(num_records)],
    'product_id': [f'{np.random.randint(1, 1000000000)}' for _ in range(num_records)]
}


df = pd.DataFrame(data)

# Save to CSV
df.to_csv('user_data.csv', index=False)
