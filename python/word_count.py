import os
import sys
from threading import Thread, Lock, Barrier

count = 0

count_lock = Lock()

def wc(content):
    return len(content.split())

def wc_file(filename):
    try:
        with open(filename, 'r', encoding='latin-1') as f:
            file_content = f.read()
        return wc(file_content)
    except FileNotFoundError:
        return 0

def wc_dir(dir_path, barrier):
    global count
    for filename in os.listdir(dir_path):
        filepath = os.path.join(dir_path, filename)
        if os.path.isfile(filepath):
            with count_lock: 
                count += wc_file(filepath)
        elif os.path.isdir(filepath):
            count += wc_dir(filepath) 
    barrier.wait()

def main():
    dirs = [os.path.abspath(sys.argv[1]), os.path.abspath(sys.argv[2]), os.path.abspath(sys.argv[3]), os.path.abspath(sys.argv[4])]
    threads = []
    barrier = Barrier(len(dirs) + 1)
    for i in range(len(dirs)):
        t = Thread(target=wc_dir, args=(dirs[i], barrier))
        threads.append(t)
        t.start()
    barrier.wait()

    
        

if __name__ == "__main__":
    main()
