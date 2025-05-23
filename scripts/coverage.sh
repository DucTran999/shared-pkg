#!/bin/bash
green() {
  echo -e "\033[32m$1\033[0m"
}

red() {
  echo -e "\033[31m$1\033[0m"
}

cyan() {
  echo -e "\033[36m$1\033[0m"
}

COVERAGE_PKG_PATH=test/coverage

cyan "🔍 Code coverage analyzing..."
echo "----------------------------------------------------------------------------------"
go test -cover ./logger/... ./scrypto/... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
echo "----------------------------------------------------------------------------------"

total_coverage=$(go tool cover -func=$COVERAGE_PKG_PATH/coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
coverage_threshold=80.0
comparison=$(echo "$total_coverage >= $coverage_threshold" | bc -l)
if [ "$comparison" -eq 0 ]; then
  red "📈 Total coverage: $total_coverage%"
  red "❌ Code coverage $total_coverage% is below the threshold of $coverage_threshold%."
  exit 1
else
  green "📈 Total coverage: $total_coverage%"
  green "✅ Code coverage $total_coverage% meets the threshold of $coverage_threshold%."
fi
