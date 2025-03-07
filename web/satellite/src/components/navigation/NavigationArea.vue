// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template src="./navigationArea.html"></template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import ProjectSelectionArea from '@/components/header/projectSelection/ProjectSelectionArea.vue';

import DocsIcon from '@/../static/images/navigation/docs.svg';
import LogoIcon from '@/../static/images/navigation/logo.svg';
import LogoTextIcon from '@/../static/images/navigation/logoText.svg';
import SupportIcon from '@/../static/images/navigation/support.svg';

import { RouteConfig } from '@/router';
import { NavigationLink } from '@/types/navigation';

@Component({
    components: {
        ProjectSelectionArea,
        LogoIcon,
        LogoTextIcon,
        DocsIcon,
        SupportIcon,
    },
})
export default class NavigationArea extends Vue {
    public areResourceItemsShown: boolean = true;
    public isResourceButtonShown: boolean = false;
    public areAccountItemsShown: boolean = true;
    public isAccountButtonShown: boolean = false;

    public toggleResourceItemsVisibility(): void {
        this.areResourceItemsShown = !this.areResourceItemsShown;
    }

    public toggleResourceButtonVisibility(): void {
        this.isResourceButtonShown = !this.isResourceButtonShown;
    }

    public toggleAccountItemsVisibility(): void {
        this.areAccountItemsShown = !this.areAccountItemsShown;
    }

    public toggleAccountButtonVisibility(): void {
        this.isAccountButtonShown = !this.isAccountButtonShown;
    }

    // TODO: Use SvgLoaderComponent to reduce markup lines
    public readonly navigation: NavigationLink[] = [
        RouteConfig.ProjectOverview.withIcon(
            `<svg class="svg" width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path class="navigation-svg-path" d="M22.073 5.3913C21.8382 5.15652 21.5252 5 21.2121 5H6.89038C6.26429 5 5.71647 5.54783 5.79473 6.09565L2.19473 6.01739C1.80342 6.09565 1.49038 6.17391 1.2556 6.48696C1.09908 6.72174 0.942555 7.03478 1.02082 7.42609L3.13386 18.4609C3.21212 19.0087 3.68169 19.4 4.22951 19.4H19.0991C19.6469 19.4 20.1165 19.0087 20.1947 18.4609L22.3078 6.33043C22.386 6.01739 22.3078 5.62609 22.073 5.3913ZM19.6469 14.7826L18.7078 9.38261C18.6295 8.75652 18.0034 8.28696 17.3773 8.28696H8.4556C8.37734 8.28696 8.29908 8.2087 8.29908 8.13043L8.14256 7.34783C8.14256 6.8 7.75125 6.33043 7.28169 6.17391L21.2121 6.09565L19.6469 14.7826Z" fill="#363840"/>
            </svg>`),
        RouteConfig.Team.withIcon(
            `<svg class="svg" width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path class="navigation-svg-path" d="M19.6621 11.1388C19.8233 10.8971 19.9039 10.6553 19.9845 10.333C20.5485 10.0913 20.8709 9.52718 20.9515 8.72136C21.032 7.99612 20.7097 7.43204 20.3068 7.10971C20.2262 5.01456 19.5816 3 16.6 3C13.6184 3 12.9738 5.01456 12.8126 7.10971C12.4097 7.43204 12.0874 8.0767 12.168 8.72136C12.2485 9.4466 12.6515 10.0913 13.135 10.333C13.2155 10.5748 13.3767 10.8971 13.4573 11.1388C12.9738 11.3 12.4903 11.4612 12.0874 11.6223C11.6845 11.4612 11.201 11.2194 10.7175 11.1388C10.8786 10.8971 10.9592 10.6553 11.0398 10.333C11.6039 10.0913 11.9262 9.52718 12.0068 8.72136C12.0874 7.99612 11.765 7.43204 11.3621 7.10971C11.201 5.01456 10.6369 3 7.57476 3C4.5932 3 3.94854 5.01456 3.78738 7.10971C3.38447 7.43204 3.06214 8.0767 3.14272 8.72136C3.2233 9.4466 3.62621 10.0913 4.10971 10.333C4.19029 10.5748 4.35146 10.8971 4.43204 11.1388C1.53107 11.9447 0 13.6369 0 16.2155C0 16.4573 0.161165 16.6184 0.32233 16.699C2.09515 17.5049 4.75437 17.9883 7.49418 17.9883C9.10583 17.9883 10.6369 17.8272 12.0068 17.5049C13.3767 17.8272 14.9078 17.9883 16.5194 17.9883C19.2592 17.9883 21.9184 17.5049 23.6913 16.699C23.8524 16.6184 24.0136 16.3767 24.0136 16.2155C24.0942 13.6369 22.5631 11.9447 19.6621 11.1388Z" fill="#354049"/>
                <path class="navigation-svg-path" d="M13.7532 12.8937C13.1891 13.1354 12.5445 13.2966 11.8192 13.4577C11.094 13.2966 10.4493 13.1354 9.88525 12.8937C9.96584 11.6043 10.6105 10.5568 11.8192 9.91211C13.1086 10.5568 13.6726 11.5238 13.7532 12.8937Z" fill="#2683FF"/>
                <path class="white" d="M14.3174 13.832C17.0572 14.4767 19.0717 15.8466 19.0717 18.7476C17.1378 19.634 14.3174 20.0369 11.8193 20.0369C9.32126 20.0369 6.42029 19.634 4.56689 18.7476C4.56689 15.8466 6.50087 14.4767 9.32126 13.832" fill="white"/>
                <path class="navigation-svg-path" d="M11.7389 20.6001C8.8379 20.6001 6.01751 20.1166 4.24469 19.2302C4.08353 19.2302 3.92236 18.9884 3.92236 18.8273C3.92236 16.0069 5.6146 14.1535 9.07964 13.3477L9.32139 14.3952C6.58159 15.0399 5.21168 16.3292 5.05052 18.4244C6.74275 19.1496 9.24081 19.5525 11.8195 19.5525C14.3981 19.5525 16.8156 19.1496 18.5884 18.4244C18.4272 16.3292 17.0573 15.0399 14.3175 14.3952L14.5593 13.3477C17.9437 14.1535 19.7165 16.0069 19.7165 18.8273C19.7165 19.069 19.5554 19.2302 19.3942 19.3108C17.4602 20.1166 14.6398 20.6001 11.7389 20.6001Z" fill="#354049"/>
                <path class="white" d="M11.7388 15.3628C10.5301 15.3628 9.16018 13.9929 8.5961 12.1395C8.11261 12.1395 7.79028 11.5754 7.7097 10.9308C7.62911 10.2861 7.95144 9.72205 8.35436 9.56088C8.43494 6.98224 9.16018 5.45117 11.7388 5.45117C14.3175 5.45117 15.0427 6.98224 15.1233 9.56088C15.5262 9.64146 15.8485 10.2055 15.7679 10.8502C15.6874 11.4949 15.365 12.0589 14.8815 12.0589C14.398 13.9929 12.9476 15.3628 11.7388 15.3628Z" fill="white"/>
                <path class="navigation-svg-path" d="M11.7387 15.926C10.2882 15.926 8.83776 14.5561 8.11251 12.6221C7.54844 12.3804 7.14552 11.7357 7.06494 11.0105C7.06494 10.2047 7.38727 9.56 7.79018 9.23767C7.95135 7.06194 8.59601 4.9668 11.7387 4.9668C14.8814 4.9668 15.5261 7.06194 15.6873 9.23767C16.1708 9.56 16.4125 10.2047 16.3319 10.9299C16.2513 11.7357 15.8484 12.3804 15.2844 12.5416C14.7203 14.4755 13.1892 15.926 11.7387 15.926ZM11.7387 6.01437C9.72416 6.01437 8.99892 6.98136 8.83776 9.56C8.83776 9.80175 8.67659 10.0435 8.43484 10.0435C8.35426 10.0435 8.1931 10.3658 8.1931 10.7687C8.1931 11.2522 8.43484 11.494 8.51543 11.494C8.75717 11.494 8.99892 11.6551 8.99892 11.8969C9.48242 13.6697 10.7717 14.7173 11.6581 14.7173C12.5446 14.7173 13.7533 13.6697 14.3174 11.8969C14.398 11.6551 14.5591 11.494 14.8009 11.494C14.8814 11.494 15.1232 11.2522 15.1232 10.7687C15.1232 10.2852 14.962 10.0435 14.8814 10.0435C14.6397 9.96291 14.4785 9.80175 14.4785 9.56C14.4785 6.98136 13.7533 6.01437 11.7387 6.01437Z" fill="#354049"/>
            </svg>`),
        RouteConfig.ApiKeys.withIcon(
            `<svg class="svg" width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path class="navigation-svg-path" d="M18.6059 9.10806C18.6059 8.85201 18.6912 8.51062 18.6912 8.25458C18.6912 4.24322 15.3626 1 11.3513 1C7.33993 1 4.0967 4.24322 4.0967 8.25458C4.0967 8.51062 4.0967 8.85201 4.18205 9.10806C1.87766 9.27875 0 11.1564 0 13.5462C0 16.0212 1.963 17.9842 4.4381 17.9842H7.76667V21.9103L6.82784 21.0568C6.57179 20.8007 6.14506 20.8007 5.88901 21.0568C5.63297 21.3128 5.63297 21.7396 5.88901 21.9956L8.27875 24.2147L10.6685 21.9956C10.9245 21.7396 10.9245 21.3128 10.6685 21.0568C10.4125 20.8007 9.98571 20.8007 9.72967 21.0568L8.79084 21.9103V17.9842H13.3143V12.0952L12.3755 12.9487C12.1194 13.2048 11.6927 13.2048 11.4366 12.9487C11.1806 12.6927 11.1806 12.2659 11.4366 12.0099L13.8264 9.79084L16.2161 12.0099C16.4722 12.2659 16.4722 12.6927 16.2161 12.9487C16.1308 13.1194 15.8747 13.1194 15.704 13.1194C15.5333 13.1194 15.3626 13.0341 15.2773 12.9487L14.3385 12.0952V17.9842H17.8377C17.9231 17.9842 18.0084 17.9842 18.0084 17.8989C20.2275 17.6429 22.0198 15.7652 22.0198 13.4608C22.6172 11.2418 20.8249 9.3641 18.6059 9.10806Z" fill="#354049"/>
            </svg>`),
        RouteConfig.Buckets.withIcon(
            `<svg class="svg" width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
               <path class="navigation-svg-path" d="M1 12.4548V20.273C1 21.273 1.81116 22.0003 2.71245 22.0003H20.1974C21.1888 22.0003 21.9099 21.1821 21.9099 20.273V12.4548C21.9099 11.4548 21.0987 10.7275 20.1974 10.7275H2.71245C1.81116 10.7275 1 11.4548 1 12.4548ZM14.97 14.0912H7.75966C7.30901 14.0912 6.85837 13.7275 6.85837 13.1821C6.85837 12.7275 7.21888 12.273 7.75966 12.273H14.97C15.4206 12.273 15.8712 12.6366 15.8712 13.1821C15.8712 13.7275 15.5107 14.0912 14.97 14.0912Z" fill="#354049"/>
               <path class="navigation-svg-path" d="M2.53227 9.81792C2.17175 9.81792 1.90137 9.54519 1.90137 9.18155V5.90882C1.90137 5.54519 2.17175 5.27246 2.53227 5.27246H20.4679C20.8284 5.27246 21.0988 5.54519 21.0988 5.90882V8.99973C21.0988 9.36337 20.8284 9.6361 20.4679 9.6361C20.1074 9.6361 19.837 9.36337 19.837 8.99973V6.54519H3.16317V9.18155C3.16317 9.54519 2.89278 9.81792 2.53227 9.81792Z" fill="#354049"/>
               <path class="navigation-svg-path" d="M20.4679 4.27273H2.53227C2.17175 4.27273 1.90137 4 1.90137 3.63636C1.90137 3.27273 2.17175 3 2.53227 3H20.4679C20.8284 3 21.0988 3.27273 21.0988 3.63636C21.0988 4 20.8284 4.27273 20.4679 4.27273Z" fill="#354049"/>
            </svg>`),
    ];

    public readonly accountNavigation: NavigationLink[] = [
        RouteConfig.Account.with(RouteConfig.Profile),
        RouteConfig.Account.with(RouteConfig.Billing),
    ];

    public onLogoClick(): void {
        this.$router.push(RouteConfig.ProjectOverview.path);
    }

    public get isProjectNotSelected(): boolean {
        return this.$store.state.projectsModule.selectedProject.id === '';
    }

    public get isResourcesDisplayingButtonShown(): boolean {
        return !this.areResourceItemsShown && this.isResourceButtonShown;
    }

    public get isResourcesHidingButtonShown(): boolean {
        return this.areResourceItemsShown && this.isResourceButtonShown;
    }

    public get isAccountItemsDisplayingButtonShown(): boolean {
        return !this.areAccountItemsShown && this.isAccountButtonShown;
    }

    public get isAccountItemsHidingButtonShown(): boolean {
        return this.areAccountItemsShown && this.isAccountButtonShown;
    }
}
</script>

<style src="./navigationArea.scss" lang="scss"></style>
