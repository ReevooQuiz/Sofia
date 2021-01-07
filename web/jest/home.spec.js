import { shallowMount } from '@vue/test-utils';
import Home from '../src/views/Home.vue';

describe('Home.vue', () => {
  it('should render correct contents', () => {
    const wrapper = shallowMount(Home);
    // expect(wrapper.find('.hello h1').text())
    //   .toEqual('Welcome to Your Vue.js App');
  });
});