### Instructions to Run

1. **Build the application:**
   ```
   go build fetch-receipt-api
   ```

2. **Run the application:**
   ```
   ./fetch-receipt-api
   ```

3. **Test your endpoints using tools like `curl` or Postman:**
   - Add your JSON receipt and replace the contents of morning.json file or manually send json data via curl
   - To process a receipt:
     ```
     curl -X POST -H "Content-Type: application/json" -d @morning.json http://localhost:8080/receipts/process
     ```

  - To get points for a receipt:
     ```
     curl http://localhost:8080/receipts/{id}/points
     ```



Left a few comments to help with my thought process through the challenge. Things I should implement for the future, automated unit testing and batch processing.
