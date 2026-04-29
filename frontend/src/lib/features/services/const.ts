// AI summary: Feature-scoped extraction of route query/mutation contracts to keep page runes and form logic focused on UI orchestration.
import type { GitProviderOption } from './type';

export const serviceTypes = [
	{ value: 'app' as const, label: 'App Service' },
	{ value: 'psql' as const, label: 'PostgreSQL Service' }
];

export const gitProviders: GitProviderOption[] = [
	{
		key: 'github',
		name: 'Github',
		icon: 'meteor-icons:github',
		api: '/provider/github/repo/list'
	},
	{
		key: 'gitlab',
		name: 'GitLab',
		icon: 'material-icon-theme:gitlab',
		api: ''
	},
	{
		key: 'bitbucket',
		name: 'BitBucket',
		icon: 'material-icon-theme:bitbucket',
		api: ''
	}
];
