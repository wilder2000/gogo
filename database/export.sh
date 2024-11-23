#!/bin/bash
 
DB_HOST="127.0.0.1"   # MySQL服务器地址
DB_USERNAME="root"    # MySQL登录用户名
DB_PASSWORD="Killwilder@2024"        # MySQL登录密码（如果有）
DATABASE="yhdb" # 要导出的数据库名称
TABLES=("s_urlmappings" "s_role" "s_operators" "s_items" "s_resources" "s_rolegroup" "s_departments" "s_groupuser" "s_roleoperator" "s_group" "s_depusers" "s_users") # 要导出的表名列表
OUTPUT_FILE="./output.sql" # 输出文件路径及名称
# 构建 mysqldump 命令
CMD="mysqldump -h $DB_HOST -u $DB_USERNAME --password=$DB_PASSWORD -t $DATABASE ${TABLES[@]} > $OUTPUT_FILE"
echo "$CMD"
eval $CMD
