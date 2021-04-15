<?php
//146. LRU 缓存机制
class LRUCache {
    public $lruQueue = [];

    /**
     * @param Integer $capacity
     */
    function __construct($capacity) {
        $this->capacity = $capacity;
    }

    /**
     * @param Integer $key
     * @return Integer
     */
    function get($key) {
        if (empty($key) || !isset($this->lruQueue[$key])) {
            return -1;
        }

        $val = $this->lruQueue[$key];
        unset($this->lruQueue[$key]);
        $this->pushToLruQueue([$key => $val]);

        return $val;
    }

    /**
     * @param Integer $key
     * @param Integer $value
     * @return NULL
     */
    function put($key, $value) {
        if (empty($this->lruQueue)) {
            $this->lruQueue[$key] = $value;
            return true;
        }


        if (isset($this->lruQueue[$key])) {
            unset($this->lruQueue[$key]);
            $this->pushToLruQueue([$key => $value]);
            return true;
        }

        if (count($this->lruQueue) < $this->capacity) {
            $this->pushToLruQueue([$key => $value]);
            return true;
        }

        array_pop($this->lruQueue);
        $this->pushToLruQueue([$key => $value]);
        return true;
    }

    private function pushToLruQueue($array)
    {
        $this->lruQueue = $array + $this->lruQueue;
    }
}