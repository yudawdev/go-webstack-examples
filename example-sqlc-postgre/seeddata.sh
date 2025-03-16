#!/bin/bash
# 基础 URL
URL="http://localhost:8080/api/v1/order/create"

# 股票符号列表
SYMBOLS=("AAPL" "MSFT" "GOOG" "AMZN" "META" "TSLA" "NVDA" "JPM" "V" "WMT")

# 订单类型
TYPES=("market" "limit")

# 进行100次请求
for i in {1..100}; do
  # 随机选择股票符号
  SYMBOL=${SYMBOLS[$((RANDOM % 10))]}

  # 随机生成数量 (1.0-50.0)
  QUANTITY=$(echo "scale=1; $RANDOM % 490 / 10 + 1" | bc)

  # 随机选择订单类型
  TYPE=${TYPES[$((RANDOM % 2))]}

  # 使用 Mac 的 uuidgen 命令生成有效 UUID
  UUID=$(uuidgen | tr '[:upper:]' '[:lower:]')

  echo "发送请求 #$i: 股票=$SYMBOL, 数量=$QUANTITY, 类型=$TYPE, UUID=$UUID"

  # 发送请求
  curl -s -X POST "$URL" \
    -H "Content-Type: application/json" \
    -d "{
      \"account_id\": \"$UUID\",
      \"symbol\": \"$SYMBOL\",
      \"quantity\": \"$QUANTITY\",
      \"type\": \"$TYPE\"
    }"

  echo -e "\n"

  # 随机暂停1-3秒
  SLEEP_TIME=$(( (RANDOM % 3) + 1 ))
  echo "等待 $SLEEP_TIME 秒..."
  sleep $SLEEP_TIME
done