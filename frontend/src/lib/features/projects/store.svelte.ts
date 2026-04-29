// AI summary: Class-based context store for project-page UI state consumed by project feature hooks.
import { getContext, setContext } from 'svelte';

interface ProjectsFeatureState {
	createDialogOpen: boolean;
	projectName: string;
	projectDescription: string;
	deletingProjectId: string;
}

class ProjectsFeatureStateClass implements ProjectsFeatureState {
	createDialogOpen = $state(false);
	projectName = $state('');
	projectDescription = $state('');
	deletingProjectId = $state('');
}

const DEFAULT_KEY = 'projects:feature:state';

export const getProjectsFeatureState = (key: string = DEFAULT_KEY) => {
	return getContext<ProjectsFeatureState>(key);
};

export const setProjectsFeatureState = (key: string = DEFAULT_KEY) => {
	return setContext(key, new ProjectsFeatureStateClass());
};
