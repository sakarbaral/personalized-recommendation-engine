from flask import Flask, request, jsonify
import joblib

app = Flask(__name__)

# Load the model
model = joblib.load('recommendation_model.pkl')

@app.route('/recommend', methods=['POST'])
def recommend():
    data = request.get_json()
    user_id = data['user_id']
    product_id = data['product_id']
    
    # Generate recommendation
    prediction = model.predict(uid=user_id, iid=product_id)
    estimated_rating = prediction.est
    
    return jsonify({'recommendation': estimated_rating})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
