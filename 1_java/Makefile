build:
	mvn clean package

assemble: build
	mkdir -p solution
	cp -r product/api/target/*.jar solution
	cp -r product/impl/target/*.jar solution
	cp -r project/target/*.jar solution

run: assemble
	cd solution && java -cp java_plugins-product-api-1.0.jar:java_plugins-product-impl-1.0.jar:java_plugins-project-1.0-SNAPSHOT.jar \
		test.impl.Main test.custom.CustomEcho

clean:
	mvn clean
	rm -rf ./solution