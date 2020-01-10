import Vue, { VueConstructor } from "vue";

export class HttpService extends Vue {
  static install(vue: VueConstructor<Vue>, opts: any) {
    Vue.prototype.$http = new HttpService();
  }

  async get(url: string) {
    return this.fetch("GET", url);
  }

  async post(url: string, data: any = null) {
    return this.fetch("POST", url, data);
  }
  async put(url: string, data: any = null) {
    return this.fetch("PUT", url, data);
  }

  async fetch(method: string, url: string, data: any = null) {
    let headers = new Headers({ "Content-Type": "application/json" });

    /*if (this.$auth.isAuthenticated()) {
            headers.append('Authorization', this.$auth.getAuthorizationHeader());
        }*/

    const response = await fetch(url, {
      method: method,
      mode: "cors",
      headers: headers,
      redirect: "follow",
      body: data != null ? JSON.stringify(data) : null
    });

    if (response.ok) {
      return await response.json();
    } else {
      throw Error(response.statusText);
    }
  }
}

declare module "vue/types/vue" {
  interface Vue {
    $http: HttpService;
  }
}
