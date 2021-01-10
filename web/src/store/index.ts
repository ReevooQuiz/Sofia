import { InjectionKey } from 'vue';
import { createStore, useStore as baseUseStore, Store } from 'vuex';

export interface State {
    navIcon: string,
    navTarget: string,
    count: number,
    logStatus:boolean,
    admin:boolean
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    state: {
        navIcon: "",
        navTarget:"home",
        count: 0,
        logStatus:false,
        admin:false
    },
    mutations: {
        changeTarget(state,newTarget){
            state.navTarget=newTarget;
        },
        changeIcon(state,newIcon){
            state.navIcon=newIcon;
        },
        changeLogStatus(state,newStatus){
            state.logStatus=newStatus;
        },
        changeAdmin(state,newAdmin){
            state.admin=newAdmin;
        }
    }
})

// define your own `useStore` composition function
export function useStore () {
    return baseUseStore(key)
}
