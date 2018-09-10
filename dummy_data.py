#!/usr/bin/env python
# -*- coding: utf-8 -*-

import psycopg2
from psycopg2.extras import Json
import uuid
import random
import string


def generate_dummy_data():
    MAX_REGISTERS = 100
    host = "localhost"
    dbname = "nuveo"
    user = "nuveo"
    password = "nuveo"
    dns = f"host={host} dbname={dbname} user={user} password={password}"
    conn = psycopg2.connect(dns)
    cur = conn.cursor()
    steps = []
    data = {}
    try:

        cur.execute("DELETE FROM nuveo.workflow")
        conn.commit()
        for i in range(0, MAX_REGISTERS):

            steps = []
            data = {}
            for j in range(0, i+1):
                steps.append("STEP " + str(j))

            data["name"] = ''.join(random.choices(string.ascii_uppercase +
                                                  string.digits,
                                                  k=10)
                                   )
            data["description"] = ''.join(random
                                          .choices(string.ascii_uppercase +
                                                   string.digits,
                                                   k=100))

            cur.execute('''INSERT INTO nuveo.workflow(uuid,status,data,steps)
                        VALUES(%s, %s, %s, %s)''',
                        (str(uuid.uuid4()), 0, Json(data), steps)
                        )
            conn.commit()
        print(f"Dados inseridos com sucesso!. " +
              f"Total de {MAX_REGISTERS} registros inseridos.")

    except (psycopg2.Error) as e:
        print(e)
    finally:
        conn.close()


if __name__ == '__main__':
    generate_dummy_data()
