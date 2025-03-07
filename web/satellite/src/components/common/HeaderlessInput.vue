// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="input-wrap">
        <div class="label-container">
            <ErrorIcon v-if="error"/>
            <h3 class="label-container__label" v-if="isLabelShown" :style="style.labelStyle">{{label}}</h3>
            <h3 class="label-container__error" v-if="error" :style="style.errorStyle">{{error}}</h3>
        </div>
        <input
            class="headerless-input"
            :class="{'inputError' : error, 'password': isPassword}"
            @input="onInput"
            @change="onInput"
            v-model="value"
            :placeholder="placeholder"
            :type="type"
            :style="style.inputStyle"
        />
        <!--2 conditions of eye image (crossed or not) -->
        <PasswordHiddenIcon
            class="input-wrap__image"
            v-if="isPasswordHiddenState"
            @click="changeVision"
        />
        <PasswordShownIcon
            class="input-wrap__image"
            v-if="isPasswordShownState"
            @click="changeVision"
        />
        <!-- end of image-->
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

import PasswordHiddenIcon from '@/../static/images/common/passwordHidden.svg';
import PasswordShownIcon from '@/../static/images/common/passwordShown.svg';
import ErrorIcon from '@/../static/images/register/ErrorInfo.svg';

// Custom input component for login page
@Component({
    components: {
        ErrorIcon,
        PasswordHiddenIcon,
        PasswordShownIcon,
    },
})
export default class HeaderlessInput extends Vue {
    private readonly textType: string = 'text';
    private readonly passwordType: string = 'password';

    private type: string = this.textType;
    private isPasswordShown: boolean = false;

    protected value: string = '';

    @Prop({default: ''})
    protected readonly label: string;
    @Prop({default: 'default'})
    protected readonly placeholder: string;
    @Prop({default: false})
    protected readonly isPassword: boolean;
    @Prop({default: '48px'})
    protected readonly height: string;
    @Prop({default: '100%'})
    protected readonly width: string;
    @Prop({default: ''})
    protected readonly error: string;
    @Prop({default: Number.MAX_SAFE_INTEGER})
    protected readonly maxSymbols: number;

    @Prop({default: false})
    private readonly isWhite: boolean;

    public constructor() {
        super();

        this.type = this.isPassword ? this.passwordType : this.textType;
    }

    // Used to set default value from parent component
    public setValue(value: string): void {
        this.value = value;
    }

    // triggers on input
    public onInput({ target }): void {
        if (target.value.length > this.maxSymbols) {
            this.value = target.value.slice(0, this.maxSymbols);
        } else {
            this.value = target.value;
        }

        this.$emit('setData', this.value);
    }

    public changeVision(): void {
        this.isPasswordShown = !this.isPasswordShown;
        if (this.isPasswordShown) {
            this.type = this.textType;

            return;
        }

        this.type = this.passwordType;
    }

    public get isLabelShown(): boolean {
        return !!(!this.error && this.label);
    }

    public get isPasswordHiddenState(): boolean {
        return this.isPassword && !this.isPasswordShown;
    }

    public get isPasswordShownState(): boolean {
        return this.isPassword && this.isPasswordShown;
    }

    protected get style(): object {
        return {
            inputStyle: {
                width: this.width,
                height: this.height,
            },
            labelStyle: {
                color: this.isWhite ? 'white' : '#354049',
            },
            errorStyle: {
                color: this.isWhite ? 'white' : '#FF5560',
            },
        };
    }
}
</script>

<style scoped lang="scss">
    .input-wrap {
        position: relative;
        width: 100%;
        font-family: 'font_regular', sans-serif;

        &__image {
            position: absolute;
            right: 25px;
            bottom: 5px;
            transform: translateY(-50%);
            z-index: 20;
            cursor: pointer;

            &:hover .input-wrap__image__path {
                fill: #2683ff !important;
            }
        }
    }

    .label-container {
        display: flex;
        justify-content: flex-start;
        align-items: flex-end;
        padding-bottom: 8px;
        flex-direction: row;

        &__label {
            font-size: 16px;
            line-height: 21px;
            color: #354049;
            margin-bottom: 0;
        }

        &__add-label {
            margin-left: 5px;
            font-size: 16px;
            line-height: 21px;
            color: rgba(56, 75, 101, 0.4);
        }

        &__error {
            font-size: 16px;
            margin: 18px 0 0 10px;
        }
    }

    .headerless-input {
        font-size: 16px;
        line-height: 21px;
        resize: none;
        height: 46px;
        padding: 0 30px 0 0;
        width: calc(100% - 30px) !important;
        text-indent: 20px;
        border: 1px solid rgba(56, 75, 101, 0.4);
        border-radius: 6px;
    }

    .headerless-input::placeholder {
        color: #384b65;
        opacity: 0.4;
    }

    .inputError::placeholder {
        color: #eb5757;
        opacity: 0.4;
    }

    .error {
        color: #ff5560;
        margin-left: 10px;
    }

    .password {
        width: calc(100% - 75px) !important;
        padding-right: 75px;
    }
</style>
