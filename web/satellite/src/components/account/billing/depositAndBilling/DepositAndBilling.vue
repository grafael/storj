// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="deposit-and-billing-area">
        <div class="deposit-and-billing-area__header">
            <h1 class="deposit-and-billing-area__header__title">Deposit & Billing History</h1>
            <div class="button" @click="onViewAllClick">View All</div>
        </div>
        <SortingHeader/>
        <BillingItem
            v-for="item in billingHistoryItems"
            :billing-item="item"
            :key="item.id"
        />
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import BillingItem from '@/components/account/billing/depositAndBilling/BillingItem.vue';
import SortingHeader from '@/components/account/billing/depositAndBilling/SortingHeader.vue';

import { RouteConfig } from '@/router';
import { BillingHistoryItem } from '@/types/payments';

@Component({
    components: {
        BillingItem,
        SortingHeader,
    },
})
export default class DepositAndBilling extends Vue {
    public onViewAllClick(): void {
        this.$router.push(RouteConfig.Account.with(RouteConfig.BillingHistory).path);
    }

    public get billingHistoryItems(): BillingHistoryItem[] {
        return this.$store.state.paymentsModule.billingHistory.slice(0, 3);
    }
}
</script>

<style scoped lang="scss">
    h1,
    span {
        margin: 0;
        color: #354049;
    }

    .deposit-and-billing-area {
        margin-bottom: 32px;
        padding: 40px;
        background-color: #fff;
        border-radius: 8px;
        font-family: 'font_regular', sans-serif;

        &__header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 40px;
            font-family: 'font_bold', sans-serif;

            &__title {
                font-size: 32px;
                line-height: 48px;
            }

            .button {
                display: flex;
                width: 120px;
                height: 48px;
                border: 1px solid #afb7c1;
                border-radius: 8px;
                align-items: center;
                justify-content: center;
                font-size: 16px;
                color: #354049;
                cursor: pointer;

                &:hover {
                    background-color: #2683ff;
                    color: #fff;
                }
            }
        }
    }

    @media screen and (max-height: 850px) {

        .deposit-and-billing-area {
            margin-bottom: 50px;
        }
    }

    @media screen and (max-height: 650px) {

        .deposit-and-billing-area {
            margin-bottom: 75px;
        }
    }
</style>
