# This script is used to simulate some process which will print the same output
# to the stdout at regular time intervals. An example of such a process is
# vagrant rsync-auto.
from random import randint
from time import sleep

for i in range(4):
    print(f"Rsyncing folders /app /tmp /config", flush=True)
    sleep(randint(1, 6))
