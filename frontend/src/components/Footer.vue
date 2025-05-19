<template>
  <footer class="footer" id="contacts">
    <div class="footer__inner container">
      <div class="footer__contacts contacts">
        <h2 class="contacts__title">Связаться с нами</h2>
        <p class="contacts__email">
          <img src="/images/icon-email.svg" alt="" class="icon" />
          <a href="mailto:email@pawfectmatch.com">email@pawfectmatch.com</a>
        </p>
        <p class="contacts__phone">
          <img src="/images/icon-phone.svg" alt="" class="icon" />
          <a href="tel:+136745677554">+13 674 567 75 54</a>
        </p>
      </div>
      <img src="/images/footer-puppy.png" alt="" class="footer__image" />
    </div>
  </footer>
</template>

<script setup></script>

<style lang="scss">
.footer {
  padding-top: rem(41);
  background-color: #262425;
  background-image:
    url('../../images/noise_transparent.png'),
    linear-gradient(
      124deg,
      rgba(38, 36, 37, 1) 0%,
      rgba(106, 81, 63, 1) 26%,
      rgba(55, 46, 40, 1) 100%
    );
  background-position: 0 0;
  background-size:
    110px 110px,
    auto;
  background-repeat: repeat, no-repeat;

  &__inner {
    display: flex;
    flex-direction: row;
    justify-content: space-between;

    @include mobile {
      flex-direction: column;
      align-items: center;
    }
  }

  &__image {
    @include mobile {
      max-width: rem(340);
    }
  }

  &__contacts {
    padding-block: rem(20) rem(60);
  }
}
</style>
