package com.jd.state;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

/**
 * Created by weiqisong on 2015/4/25.
 */
public class MachineState {


    //
    private int expire;

    private Map<String,Map<Integer,Transaction>> state = new HashMap<>();

    private ConcurrentMap<Long,String> preStates = new ConcurrentHashMap<>();




}

class Transaction {

    private String preType;
    private String currentType;
    private boolean isEnd;

    public String getPreType() {
        return preType;
    }

    public void setPreType(String preType) {
        this.preType = preType;
    }

    public String getCurrentType() {
        return currentType;
    }

    public void setCurrentType(String currentType) {
        this.currentType = currentType;
    }
}
