#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import json
import questplus as qp


def handler(event, context):
    arg = json.load(event)
    print(arg)

    # Initialize the QUEST+ staircase.
    q = qp.QuestPlus(**arg["qp_params"])

    for row in arg["data"]:
        # エントロピーを内部で計算
        # _ = q.next_stim["intensity"]
        stim = int(row["move_width"])
        response = row["response"]
        q.update(stim=dict(intensity=float(stim / 10)), outcome=dict(response=response))

    return json.dump(dict(next_stim=q.next_stim["intensity"]))


