#include <openssl/evp.h>
#include <stdio.h>
#include <stdlib.h>
int main() {
    // Assuming you have an EVP_PKEY structure representing the SM2 private key
    EVP_PKEY *sm2_privkey = /* initialize with your SM2 private key */;
    // Encrypted data buffer and length
    unsigned char *ciphertext = /* initialize with your encrypted data */;
    size_t ciphertext_len = /* length of the encrypted data */;
    // Step 1: Call EVP_PKEY_decrypt() to get the required output buffer size
    size_t plaintext_len;
    int result = EVP_PKEY_decrypt(NULL, NULL, &plaintext_len, ciphertext, ciphertext_len, sm2_privkey);
    if (result != 1) {
        // Handle error
        return 1;
    }
    // Allocate a buffer of the required size
    unsigned char *plaintext = (unsigned char *)malloc(plaintext_len);
    if (plaintext == NULL) {
        // Handle memory allocation failure
        return 1;
    }
    // Step 2: Call EVP_PKEY_decrypt() again with the allocated buffer
    result = EVP_PKEY_decrypt(plaintext, &plaintext_len, NULL, ciphertext, ciphertext_len, sm2_privkey);
    if (result != 1) {
        // Handle error
        free(plaintext);
        return 1;
    }
    // Use the decrypted plaintext data
    printf("Decrypted plaintext: %.*s\n", (int)plaintext_len, plaintext);
    // Clean up
    free(plaintext);
    return 0;
}