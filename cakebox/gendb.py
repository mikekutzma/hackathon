import pandas as pd
import sqlite3

raw_data = pd.read_csv("birthdays.csv", sep = " ")
raw_data = raw_data.astype({"birthMonth": "Int64", "birthDay": "Int64"})

conn = sqlite3.connect("sqlite.db")

raw_data.to_sql("birthdays", conn, if_exists="replace")

print("Wrote db")
