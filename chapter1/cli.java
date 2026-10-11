public class Main {
  public static void main(String[] args) {
    ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 8080)
      .usePlaintext(true)
      .build();
    ProductInfoGrpc.ProductInfoBlockingStub stub =
      ProductInfoGrpc.newBlockingStub(channel);
    StringValue productID = stub.addProduct(
      Product.newBuilder()
      .setName("Apple iPhone 11")
      .setDescription("Meet Apple iPhone 11." +
      "All-new dual-camera system with " +
      "Ultra Wide and Night mode.")
      .build());
  }
}
