<template>
  <section class="section">
    <div class="section__usability usability container">
      <h2 class="usability__title">Почему это удобно?</h2>
      <ul class="usability__list">
        <li class="usability__item">
          <div class="number"></div>
          <p class="usability__item-title">
            Вместо того чтобы тратить время на поиск в разных приютах, вы находите всё в одном месте
          </p>
        </li>
        <li class="usability__item">
          <div class="number"></div>
          <p class="usability__item-title">
            Вместо того чтобы тратить время на поиск в разных приютах, вы находите всё в одном месте
          </p>
        </li>
        <li class="usability__item">
          <div class="number"></div>
          <p class="usability__item-title">
            Вместо того чтобы тратить время на поиск в разных приютах, вы находите всё в одном месте
          </p>
        </li>
      </ul>
    </div>
  </section>
</template>

<script setup></script>

<style lang="scss" scoped>
.usability {
  background-image: url('../../images/donation-dog.png');
  background-repeat: no-repeat;
  background-position: 0 102%;
  position: relative;
  z-index: 1;
  overflow: hidden;

  &__title {
    color: var(--color-light);
    text-align: center;
    margin-bottom: rem(115);
    font-style: italic;
    font-weight: 400;
  }

  &__list {
    display: flex;
    flex-direction: column;
    row-gap: rem(115);

    @include tablet {
      margin-left: 130px;
    }

    @include mobile-s {
      gap: rem(80);
    }
  }

  &__item {
    max-width: rem(613);

    &-title {
      color: var(--color-light);
      font-size: rem(20);
      font-style: italic;

      @include mobile {
        font-size: rem(16);
      }
    }

    &:nth-child(1) {
      margin-left: rem(185);
      position: relative;

      &::before {
        position: absolute;
        content: '1';
        color: var(--color-light);
        font-size: rem(128);
        font-style: italic;
        right: calc(100% + rem(27));
        top: rem(-70);
        z-index: 1;
      }

      &::after {
        position: absolute;
        content: '';
        background-image: url('../../images/circle.png');
        width: rem(99);
        height: rem(99);
        top: rem(5);
        right: calc(100% + rem(21));
      }

      @include tablet {
        margin-left: 0;
      }

      @include mobile {
        max-width: none;
      }
    }

    &:nth-child(2) {
      margin-left: rem(348);
      position: relative;

      &::before {
        position: absolute;
        content: '2';
        color: var(--color-light);
        font-size: rem(128);
        font-style: italic;
        right: calc(100% + rem(27));
        top: rem(-70);
        z-index: 1;
      }

      &::after {
        position: absolute;
        content: '';
        background-image: url('../../images/circle.png');
        width: rem(99);
        height: rem(99);
        top: rem(5);
        right: calc(100% + rem(21));
      }

      @include tablet {
        margin-left: 0;
      }

      @include mobile {
        max-width: none;
      }
    }

    &:nth-child(3) {
      margin-left: rem(526);
      max-width: rem(711);
      position: relative;

      &::before {
        position: absolute;
        content: '3';
        color: var(--color-light);
        font-size: rem(128);
        font-style: italic;
        right: calc(100% + rem(27));
        top: rem(-70);
        z-index: 1;
      }

      &::after {
        position: absolute;
        content: '';
        background-image: url('../../images/circle.png');
        width: rem(99);
        height: rem(99);
        top: rem(5);
        right: calc(100% + rem(21));
      }

      @include tablet {
        margin-left: 0;
      }

      @include mobile {
        max-width: none;
      }
    }
  }
}
</style>
