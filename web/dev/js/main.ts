import Alpine from 'alpinejs';
import { userId } from './shared/vars';

console.log(userId, ' Hello!');

Alpine.data('test', () => ({
	items: [],
	hello: () => {
		console.log('hello2');
	},
}));
