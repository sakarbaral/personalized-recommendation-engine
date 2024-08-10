import pandas as pd
from surprise import Dataset, Reader, SVD, accuracy
from surprise.model_selection import train_test_split
import joblib


data = pd.read_csv('user_data.csv')


event_type_mapping = {'view': 1, 'click': 2, 'purchase': 3}
data['event_type_encoded'] = data['event_type'].map(event_type_mapping)

reader = Reader(rating_scale=(1, 3))
data_surprise = Dataset.load_from_df(data[['user_id', 'product_id', 'event_type_encoded']], reader)

trainset, testset = train_test_split(data_surprise, test_size=0.25)

model = SVD()
model.fit(trainset)

predictions = model.test(testset)
accuracy.rmse(predictions)

joblib.dump(model, 'recommendation_model.pkl')
