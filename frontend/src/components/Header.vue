<template>
  <header class="header">
    <div :class="['header__inner', { 'header__inner--white-bg': isWhite }, 'container']">
      <div class="header__logo">
        <router-link to="/">
          <img :src="isWhite ? './images/logo-animals.svg' : './images/Logo-1.svg'" alt="" />
        </router-link>
      </div>
      <div :class="['header__burger', { 'header__burger--black': isWhite }]">
        <span class="header__burger-item"></span>
        <span class="header__burger-item"></span>
        <span class="header__burger-item"></span>
      </div>
      <nav class="header__nav">
        <ul :class="['header__list', { 'header__list--black': isWhite }]">
          <li class="header__item" v-for="link in navLinks" :key="link.text">
            <template v-if="link.to === 'contacts'">
              <a
                href="#contacts"
                @click.prevent="scrollToContacts"
                :class="['header__link', { 'header__link--black': isWhite }]"
              >
                {{ link.text }}
              </a>
            </template>
            <template v-else>
              <router-link
                :to="link.to"
                :class="['header__link', { 'header__link--black': isWhite }]"
              >
                {{ link.text }}
              </router-link>
            </template>
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script setup>
import { computed } from 'vue'
import { defineProps } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'default', // или 'white'
  },
})

const isWhite = computed(() => props.variant === 'white')

const navLinks = computed(() => [
  { to: isWhite.value ? '/#about' : '#about', text: 'Об агрегаторе' },
  { to: '/animals', text: 'Животные' },
  { to: '/profile', text: 'Личный кабинет' },
  { to: '#contacts', text: 'Контакты' },
])

const scrollToContacts = () => {
  const element = document.getElementById('contacts')
  if (element) {
    element.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<style lang="scss" scoped>
.header {
  position: relative;
  z-index: 1;

  &:has(.header__inner--white-bg) {
    background-color: var(--color-light);
  }

  &__inner {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: rem(30);

    @include mobile {
      flex-direction: column;
      margin-top: rem(15);
      gap: rem(20);
    }

    @include mobile-s {
      flex-direction: row;
      justify-content: space-between;
    }

    &--white-bg {
      background-color: var(--color-light);
      margin-block: rem(30);
    }
  }

  &__nav {
    @include mobile-s {
      display: none;
    }
  }

  &__list {
    display: flex;
    column-gap: rem(20);
    color: #cdcdcd;
    font-size: rem(15);
    font-weight: 400;

    &--black {
      color: var(--color-dark-2);
    }

    @include mobile-s {
      font-size: rem(12);
    }
  }

  &__logo {
    flex-shrink: 0;
    max-width: 60%;
  }

  &__link {
    position: relative;

    &::after {
      content: '';
      position: absolute;
      top: 100%;
      left: 50%;
      translate: -50%;
      width: 0;
      height: 0.125rem;
      background-color: var(--color-accent);
      transition-duration: var(--transition-duration);
    }

    @include hover {
      color: var(--color-light);

      &::after {
        width: 100%;
      }
    }

    &--black {
      @include hover {
        opacity: 1;
        color: var(--color-dark-1);
      }
    }
  }

  &__burger {
    display: flex;
    gap: rem(5);
    flex-direction: column;
    display: none;

    @include mobile-s {
      display: flex;
    }

    &-item {
      background-color: #fff;
      width: rem(30);
      height: rem(2);
    }

    &--black {
      .header__burger-item {
        background-color: #545454;
      }
    }
  }
}
</style>
