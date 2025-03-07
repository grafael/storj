// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="new-project-popup-container" @keyup.enter="createProjectClick" @keyup.esc="onCloseClick">
        <div class="new-project-popup" id="newProjectPopup" >
            <div class="new-project-popup__info-panel-container">
                <h2 class="new-project-popup__info-panel-container__main-label-text">Create a Project</h2>
                <CreateProjectIcon alt="Create project illustration"/>
            </div>
            <div class="new-project-popup__form-container">
                <HeaderedInput
                    label="Project Name"
                    additional-label="Up To 20 Characters"
                    placeholder="Enter Project Name"
                    class="full-input"
                    width="100%"
                    max-symbols="20"
                    :error="nameError"
                    @setData="setProjectName"
                />
                <HeaderedInput
                    label="Description"
                    placeholder="Enter Project Description"
                    additional-label="Optional"
                    class="full-input"
                    is-multiline="true"
                    height="100px"
                    width="100%"
                    @setData="setProjectDescription"
                />
                <div class="new-project-popup__form-container__button-container">
                    <VButton
                        label="Cancel"
                        width="205px"
                        height="48px"
                        :on-press="onCloseClick"
                        is-white="true"
                    />
                    <VButton
                        label="Next"
                        width="205px"
                        height="48px"
                        :on-press="createProjectClick"
                    />
                </div>
            </div>
            <div class="new-project-popup__close-cross-container" @click="onCloseClick">
                <CloseCrossIcon/>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import HeaderedInput from '@/components/common/HeaderedInput.vue';
import VButton from '@/components/common/VButton.vue';

import CloseCrossIcon from '@/../static/images/common/closeCross.svg';
import CreateProjectIcon from '@/../static/images/project/createProject.svg';

import { BUCKET_ACTIONS } from '@/store/modules/buckets';
import { PROJECTS_ACTIONS } from '@/store/modules/projects';
import { PROJECT_USAGE_ACTIONS } from '@/store/modules/usage';
import { CreateProjectModel, Project } from '@/types/projects';
import {
    API_KEYS_ACTIONS,
    APP_STATE_ACTIONS,
    PM_ACTIONS,
} from '@/utils/constants/actionNames';

@Component({
    components: {
        HeaderedInput,
        VButton,
        CreateProjectIcon,
        CloseCrossIcon,
    },
})
export default class NewProjectPopup extends Vue {
    private projectName: string = '';
    private description: string = '';
    private nameError: string = '';
    private createdProjectId: string = '';
    private isLoading: boolean = false;

    public setProjectName (value: string): void {
        this.projectName = value;
        this.nameError = '';
    }

    public setProjectDescription (value: string): void {
        this.description = value;
    }

    public onCloseClick (): void {
        this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_NEW_PROJ);
    }

    public async createProjectClick (): Promise<void> {
        if (this.isLoading) {
            return;
        }

        this.isLoading = true;

        if (!this.validateProjectName()) {
            this.isLoading = false;

            return;
        }

        try {
            const project = await this.createProject();
            this.createdProjectId = project.id;
        } catch (e) {
            this.isLoading = false;
            await this.$notify.error(e.message);
            this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_NEW_PROJ);

            return;
        }

        this.selectCreatedProject();

        try {
            await this.fetchProjectMembers();
        } catch (e) {
            await this.$notify.error(e.message);
        }

        this.clearApiKeys();

        this.clearUsage();

        this.clearBucketUsage();

        this.checkIfsFirstProject();

        this.isLoading = false;
    }

    private validateProjectName(): boolean {
        this.projectName = this.projectName.trim();

        const rgx = /^[^/]+$/;
        if (!rgx.test(this.projectName)) {
            this.nameError = 'Name for project is invalid!';

            return false;
        }

        if (this.projectName.length > 20) {
            this.nameError = 'Name should be less than 21 character!';

            return false;
        }

        return true;
    }

    private async createProject(): Promise<Project> {
        const project: CreateProjectModel = {
            name: this.projectName,
            description: this.description,
        };

        return await this.$store.dispatch(PROJECTS_ACTIONS.CREATE, project);
    }

    private selectCreatedProject(): void {
        this.$store.dispatch(PROJECTS_ACTIONS.SELECT, this.createdProjectId);

        this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_NEW_PROJ);
    }

    private checkIfsFirstProject(): void {
        const isFirstProject = this.$store.state.projectsModule.projects.length === 1;

        isFirstProject
            ? this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_SUCCESSFUL_PROJECT_CREATION_POPUP)
            : this.notifySuccess('Project created successfully!');
    }

    private async fetchProjectMembers(): Promise<void> {
        await this.$store.dispatch(PM_ACTIONS.CLEAR);
        const fistPage = 1;
        await this.$store.dispatch(PM_ACTIONS.FETCH, fistPage);
    }

    private clearApiKeys(): void {
        this.$store.dispatch(API_KEYS_ACTIONS.CLEAR);
    }

    private clearUsage(): void {
        this.$store.dispatch(PROJECT_USAGE_ACTIONS.CLEAR);
    }

    private clearBucketUsage(): void {
        this.$store.dispatch(BUCKET_ACTIONS.SET_SEARCH, '');
        this.$store.dispatch(BUCKET_ACTIONS.CLEAR);
    }

    private async notifySuccess(message: string): Promise<void> {
        await this.$notify.success(message);
    }
}
</script>

<style scoped lang="scss">
    .new-project-popup-container {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(134, 134, 148, 0.4);
        z-index: 1121;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .input-container.full-input {
        width: 100%;
    }

    .new-project-popup {
        width: 100%;
        max-width: 845px;
        height: 400px;
        background-color: #fff;
        border-radius: 6px;
        display: flex;
        flex-direction: row;
        align-items: center;
        position: relative;
        justify-content: center;
        padding: 100px 100px 100px 80px;

        &__info-panel-container {
            display: flex;
            flex-direction: column;
            justify-content: flex-start;
            align-items: center;
            margin-right: 55px;
            height: 535px;

            &__main-label-text {
                font-family: 'font_bold', sans-serif;
                font-size: 32px;
                line-height: 39px;
                color: #384b65;
                margin-bottom: 60px;
                margin-top: 50px;
            }
        }

        &__form-container {
            width: 100%;
            max-width: 520px;

            &__button-container {
                width: 100%;
                display: flex;
                flex-direction: row;
                justify-content: space-between;
                align-items: center;
                margin-top: 30px;
            }
        }

        &__close-cross-container {
            display: flex;
            justify-content: center;
            align-items: center;
            position: absolute;
            right: 30px;
            top: 40px;
            height: 24px;
            width: 24px;
            cursor: pointer;

            &:hover .close-cross-svg-path {
                fill: #2683ff;
            }
        }
    }

    @media screen and (max-width: 720px) {

        .new-project-popup {

            &__info-panel-container {
                display: none;
            }

            &__form-container {

                &__button-container {
                    width: 100%;
                }
            }
        }
    }
</style>
