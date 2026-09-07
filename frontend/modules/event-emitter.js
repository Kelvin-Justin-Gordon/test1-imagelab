//Permits asynchronous communication so that app.js can announce when a state has been changed, all without render.js and app.js needing to know about each other 
export class EventEmitter {
    #listeners = new Map();

    on(event, handler){
        if(!this.#listeners.has(event)){
            this.#listeners.set(event, new Set());
        }
        this.#listeners.get(event).add(handler);
        return () => this.off(event, handler);
    }

    off(event, handler){
        this.#listeners.get(event)?.delete(handler);
    }

    emit(event, payload){
        this.#listeners.get(event)?.forEach((handler) => handler(payload));
    }
}